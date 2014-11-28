package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const VERSION = "0.1.0"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type FileEntry struct {
	Filename string `json:"filename"`
	IsDir    bool   `json:"directory"`
}

var (
	// Regular expression to match all supported video files
	Extensions = regexp.MustCompile(".(avi|mpg|mov|flv|wmv|asf|mpeg|m4v|divx|mp4|mkv)$")

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
	}

	// Path where all media files are stored
	MediaPath string

	// Path to omxplayer executable
	OmxPath string

	// Child process for spawning omxplayer
	Omx *exec.Cmd

	// Child process STDIN pipe to send commands
	OmxIn io.WriteCloser

	// Channel to pass along commands to the player routine
	Command chan string
)

// Returns true if specified file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Scan given path for all directories and matching video files.
// If nothing was found it will return an empty slice.
func scanPath(path string) []FileEntry {
	entries := make([]FileEntry, 0)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return entries
	}

	for _, file := range files {
		entry := FileEntry{
			Filename: file.Name(),
			IsDir:    file.IsDir(),
		}

		// Do not include non-video files in the list
		if !file.IsDir() && !omxCanPlay(file.Name()) {
			continue
		}

		entries = append(entries, entry)
	}

	return entries
}

// Determine the full path to omxplayer executable. Returns error if not found.
func omxDetect() error {
	buff, err := exec.Command("which", "omxplayer").Output()
	if err != nil {
		return err
	}

	// Set path in global variable
	OmxPath = strings.TrimSpace(string(buff))

	return nil
}

// Start command listener. Commands are coming in through a channel.
func omxListen() {
	Command = make(chan string)

	for {
		command := <-Command

		// Skip command handling of omx player is not active
		if Omx == nil {
			continue
		}

		// Send command to the player
		omxWrite(command)

		// Attempt to kill the process if stop command is requested
		if command == "stop" {
			Omx.Process.Kill()
		}
	}
}

// Start omxplayer playback for a given video file. Returns error if start fails.
func omxPlay(file string) error {
	Omx = exec.Command(
		OmxPath,     // path to omxplayer executable
		"--refresh", // adjust framerate/resolution to video
		"--blank",   // set background to black
		"--adev",    // audio out device
		"hdmi",      // using hdmi for audio/video
		file,        // path to video file
	)

	// Grab child process STDIN
	stdin, err := Omx.StdinPipe()
	if err != nil {
		return err
	}

	defer stdin.Close()

	// Redirect output for debugging purposes
	Omx.Stdout = os.Stdout

	// Start omxplayer execution.
	// If successful, something will appear on HDMI display.
	err = Omx.Start()
	if err != nil {
		return err
	}

	// Make child's STDIN globally available
	OmxIn = stdin

	// Wait until child process is finished
	err = Omx.Wait()
	if err != nil {
		fmt.Fprintln(os.Stdout, "Process exited with error:", err)
	}

	omxCleanup()

	return nil
}

// Write a command string to the omxplayer process's STDIN
func omxWrite(command string) {
	if OmxIn != nil {
		io.WriteString(OmxIn, Commands[command])
	}
}

// Terminate any running omxplayer processes. Fixes random hangs.
func omxKill() {
	exec.Command("killall", "omxplayer.bin").Output()
	exec.Command("killall", "omxplayer").Output()
}

// Reset internal state and stop any running processes
func omxCleanup() {
	Omx = nil
	OmxIn = nil

	omxKill()
}

// Check if player is currently active
func omxIsActive() bool {
	return Omx != nil
}

// Check if player can play the file
func omxCanPlay(path string) bool {
	if Extensions.Match([]byte(path)) {
		return true
	}

	return false
}

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

	file = fmt.Sprintf("%s/%s", MediaPath, file)

	if !fileExists(file) {
		c.JSON(400, Response{false, "File does not exist"})
		return
	}

	if !omxCanPlay(file) {
		c.JSON(400, Response{false, "File cannot be played"})
		return
	}

	go omxPlay(file)

	c.JSON(200, Response{true, "OK"})
}

func httpStatus(c *gin.Context) {
	c.String(200, fmt.Sprintf(`{"running":%v}`, omxIsActive()))
}

func httpIndex(c *gin.Context) {
	data, err := Asset("static/index.html")

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Data(200, "text/html; charset=utf-8", data)
}

func terminate(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func usage() {
	terminate("Usage: omxremote path/to/media/dir", 0)
}

func main() {
	fmt.Printf("omxremote v%v\n", VERSION)

	if len(os.Args) < 2 {
		usage()
	}

	// Get path from arguments and remove trailing slash
	MediaPath = strings.TrimRight(os.Args[1], "/")

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

	// Disable debugging mode
	gin.SetMode("release")

	// Setup HTTP server
	router := gin.Default()

	router.GET("/", httpIndex)
	router.GET("/status", httpStatus)
	router.GET("/browse", httpBrowse)
	router.GET("/play", httpPlay)
	router.GET("/command/:command", httpCommand)

	fmt.Println("Starting server on 0.0.0.0:8080")
	router.Run(":8080")
}
