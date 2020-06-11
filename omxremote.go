package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const VERSION = "0.2.0"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type StatusResponse struct {
	Running  bool   `json:"running"`            // True if player is running
	File     string `json:"file"`               // Path to current media file
	Name     string `json:"name"`               // Titleized filename
	Position string `json:"position,omitempty"` // Current position in the movie
	Duration string `json:"duration,omitempty"` // Movie duration
}

type FileEntry struct {
	Filename string `json:"filename"`
	IsDir    bool   `json:"directory"`
}

var (
	// Regular expression to match all supported video files
	RegexFormats = regexp.MustCompile(`.(avi|mpg|mov|flv|wmv|asf|mpeg|m4v|divx|mp4|ogm|mkv|mp4)$`)

	// Regular expression to convert filenames to titles
	RegexBrackets = regexp.MustCompile(`[\(\[\]\)]`)
	RegexYear     = regexp.MustCompile(`((19|20)[\d]{2})`)
	RegexEpisode  = regexp.MustCompile(`(?i)S[\d]+E[\d]+`)
	RegexJunk     = regexp.MustCompile(`(?i)(1080p|720p|3d|brrip|bluray|webrip|x264|aac)`)
	RegexSpace    = regexp.MustCompile(`\s{2,}`)

	// OMXPlayer control commands, these are piped via STDIN to omxplayer process
	Commands = map[string]string{
		"pause":             "p",            // Pause/continue playback
		"stop":              "q",            // Stop playback and exit
		"volume_up":         "+",            // Change volume by +3dB
		"volume_down":       "-",            // Change volume by -3dB
		"subtitles":         "s",            // Enable/disable subtitles
		"seek_back":         "\x1b\x5b\x44", // Seek -30 seconds
		"seek_back_fast":    "\x1b\x5b\x42", // Seek -600 second
		"seek_forward":      "\x1b\x5b\x43", // Seek +30 second
		"seek_forward_fast": "\x1b\x5b\x41", // Seek +600 seconds
		"next_audiotrack":   "k",            // select next audio track
		"prev_audiotrack":   "j",            // select previous audio track
	}

	MediaPath   string         // Path where all media files are stored
	OmxPath     string         // Path to omxplayer executable
	Omx         *exec.Cmd      // Child process for spawning omxplayer
	OmxIn       io.WriteCloser // Child process STDIN pipe to send commands
	Command     chan string    // Channel to pass along commands to the player routine
	CurrentFile string         // Currently playing media file name
	Zeroconf    bool           // Enable Zeroconf discovery
	Frontend    bool           // Serve frontend app
	stream      *Stream        // Current stream
)

func httpBrowse(c *gin.Context) {
	path := c.Request.FormValue("path")

	if path != "" {
		path = fmt.Sprintf("%s/%s", MediaPath, path)
	} else {
		path = MediaPath
	}

	c.JSON(200, scanPath(path))
}

func httpCommand(c *gin.Context) {
	val := c.Params.ByName("command")

	if _, ok := Commands[val]; !ok {
		c.JSON(400, Response{false, "Invalid command"})
		return
	}

	fmt.Println("Received command:", val)

	// Handle requested commmand
	Command <- val

	c.JSON(200, Response{true, "OK"})
}

func httpServe(c *gin.Context) {
	file := c.Request.URL.Query().Get("file")
	if file == "" {
		return
	}

	file = fmt.Sprintf("%s/%s", MediaPath, file)
	if !fileExists(file) {
		c.String(404, "Not found")
		return
	}

	if !omxCanPlay(file) {
		c.String(400, "Invalid format")
		return
	}

	http.ServeFile(c.Writer, c.Request, file)
}

func httpPlay(c *gin.Context) {
	if omxIsActive() {
		c.JSON(400, Response{false, "Player is already running"})
		return
	}

	file := c.Request.FormValue("file")
	if file == "" {
		c.JSON(400, Response{false, "File is required"})
		return
	}

	if !strings.HasPrefix(file, "http") {
		file = fmt.Sprintf("%s/%s", MediaPath, file)

		if !fileExists(file) {
			c.JSON(400, Response{false, "File does not exist"})
			return
		}

		if !omxCanPlay(file) {
			c.JSON(400, Response{false, "File cannot be played"})
			return
		}
	}

	go omxPlay(file)

	c.JSON(200, Response{true, "OK"})
}

func httpInfo(c *gin.Context) {
	file := c.Request.FormValue("file")
	if file == "" {
		c.JSON(400, Response{false, "File is required"})
		return
	}

	file = fmt.Sprintf("%s/%s", MediaPath, file)
	if !fileExists(file) {
		c.JSON(400, Response{false, "File does not exist"})
		return
	}

	info, err := omxInfo(file)
	if err != nil {
		c.JSON(400, Response{false, err.Error()})
		return
	}

	c.JSON(200, info)
}

func httpRemoveFile(c *gin.Context) {
	file := strings.TrimSpace(c.Request.FormValue("file"))
	if file == "" {
		c.JSON(400, Response{false, "File is required"})
		return
	}

	fullPath := fmt.Sprintf("%s/%s", MediaPath, file)
	if !fileExists(fullPath) {
		c.JSON(400, Response{false, "File does not exist"})
		return
	}

	if err := os.RemoveAll(fullPath); err != nil {
		c.JSON(400, Response{false, err.Error()})
		return
	}

	c.JSON(200, Response{true, "OK"})
}

func httpStatus(c *gin.Context) {
	resp := StatusResponse{
		Running: omxIsActive(),
		File:    CurrentFile,
		Name:    fileToTitle(filepath.Base(CurrentFile)),
	}

	if stream != nil {
		resp.Duration = durationFromSeconds(stream.duration)
		resp.Position = stream.pos.String()
	}

	c.JSON(200, resp)
}

func httpIndex(c *gin.Context) {
	data, err := Asset("static/index.html")

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Data(200, "text/html; charset=utf-8", data)
}

// Retrieve information about the host: uptime, storage, etc
func httpHost(c *gin.Context) {
	output := map[string]string{}
	commands := map[string]string{
		"os":      "uname -a",
		"uptime":  "uptime",
		"storage": "df -m",
		"memory":  "free -m",
		"temp":    "/opt/vc/bin/vcgencmd measure_temp | egrep -o '[0-9]*\\.[0-9]*'",
	}

	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(len(commands))

	for k, v := range commands {
		go func(key, name string) {
			defer wg.Done()

			out := bytes.NewBuffer(nil)

			cmd := exec.Command("bash", "-c", name)
			cmd.Stdout = out
			cmd.Stderr = out

			if err := cmd.Run(); err != nil {
				log.Println("Failed to execute command", key, err)

				lock.Lock()
				output[key] = "N/A"
				lock.Unlock()

				return
			}

			lock.Lock()
			output[key] = strings.TrimSpace(out.String())
			lock.Unlock()

		}(k, v)
	}

	wg.Wait()

	c.JSON(200, gin.H{
		"uptime":  output["uptime"],
		"os":      output["os"],
		"storage": parseStorageInfo(output["storage"]),
		"memory":  parseMemoryInfo(output["memory"]),
		"temp":    parseTemperature(output["temp"]),
	})
}

// Reboot the operating system
// POST /reboot
func httpReboot(c *gin.Context) {
	if err := exec.Command("sudo", "reboot").Run(); err != nil {
		c.JSON(400, Response{Success: false, Message: err.Error()})
		return
	}
	c.JSON(200, Response{Success: true})
}

func terminate(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func usage() {
	terminate("Usage: omxremote path/to/media/dir", 0)
}

func init() {
	var printVersion bool

	flag.StringVar(&MediaPath, "media", "./", "Path to media files")
	flag.BoolVar(&Frontend, "frontend", true, "Enable frontend applicaiton")
	flag.BoolVar(&Zeroconf, "zeroconf", true, "Enable service advertisement with Zeroconf")
	flag.BoolVar(&printVersion, "v", false, "Print version")
	flag.Parse()

	if printVersion {
		fmt.Printf("omxremote v%v\n", VERSION)
		os.Exit(0)
	}
}

func main() {
	// Expand media path if needed
	MediaPath = strings.Replace(MediaPath, "~", os.Getenv("HOME"), 1)

	// Get path from arguments and remove trailing slash
	MediaPath = strings.TrimRight(MediaPath, "/")

	if !fileExists(MediaPath) {
		terminate(fmt.Sprintf("Directory does not exist: %s", MediaPath), 1)
	}

	// Check if player is installed
	if omxDetect() != nil {
		terminate("omxplayer is not installed", 1)
	}

	// Make sure nothing is running
	omxCleanup()

	// Start a remote command listener
	go omxListen()

	// Start zeroconf service advertisement
	if Zeroconf {
		stopZeroconf := make(chan bool)
		go startZeroConfAdvertisement(stopZeroconf)
	}

	// Disable debugging mode
	gin.SetMode("release")

	// Setup HTTP server
	router := gin.Default()

	// Handle CORS
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Expose-Headers", "*")
	})

	// Server frontend application only if its enabled
	if Frontend == true {
		router.GET("/", httpIndex)
	}

	router.GET("/status", httpStatus)
	router.GET("/browse", httpBrowse)
	router.GET("/info", httpInfo)
	router.GET("/play", httpPlay)
	router.GET("/serve", httpServe)
	router.POST("/remove", httpRemoveFile)
	router.GET("/command/:command", httpCommand)
	router.GET("/host", httpHost)
	router.POST("/reboot", httpReboot)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	fmt.Println("Starting server on " + host + ":" + port)
	router.Run(host + ":" + port)
}
