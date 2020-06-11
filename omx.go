package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type FileInfo struct {
	Duration string `json:"duration"`
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

func omxInfo(file string) (*FileInfo, error) {
	output := bytes.NewBuffer(nil)

	cmd := exec.Command(OmxPath, "--info", file)
	cmd.Stdout = output
	cmd.Stderr = output

	// Valid `omxplayer --info` exits with status 1, so we just log the error
	if err := cmd.Run(); err != nil {
		log.Println("omxplayer --info returned error:", err)
	}

	data := string(output.String())
	info := &FileInfo{}

	matches := durationRegexp.FindAllStringSubmatch(data, 1)
	if len(matches) > 0 {
		info.Duration = matches[0][1]
	}

	return info, nil
}

// Start omxplayer playback for a given video file. Returns error if start fails.
func omxPlay(file string) error {

	addnopts := strings.Split(OmxAddnOpts, " ")
	addnopts = append(addnopts, "--stats")      // print stats to stdout (buffers, time, etc)
	addnopts = append(addnopts, "--with-info")  // print stats about streams before playback
	addnopts = append(addnopts, "--refresh")    // adjust framerate/resolution to video
	addnopts = append(addnopts, "--blank")      // set background to black
	addnopts = append(addnopts, file)

	Omx = exec.Command(OmxPath, addnopts...)

	// Grab child process STDIN
	stdin, err := Omx.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	stderr, err := Omx.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	stdout, err := Omx.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	stream = NewStream()
	go stream.Start(stdout, stderr)

	// Start omxplayer execution.
	// If successful, something will appear on HDMI display.
	err = Omx.Start()
	if err != nil {
		return err
	}

	// Set current file
	CurrentFile = file

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
	if OmxIn == nil {
		log.Println("Cant write to omxplayer stdin: not setup")
		return
	}

	_, err := io.WriteString(OmxIn, Commands[command])
	if err != nil {
		log.Println("Cant write to omxplayer:", err)
	}
}

// Terminate any running omxplayer processes. Fixes random hangs.
func omxKill() {
	exec.Command("killall", "omxplayer.bin").Run()
	exec.Command("killall", "omxplayer").Run()
}

// Reset internal state and stop any running processes
func omxCleanup() {
	Omx = nil
	OmxIn = nil
	CurrentFile = ""

	omxKill()

	stream = nil
}

// Check if player is currently active
func omxIsActive() bool {
	return Omx != nil
}

// Check if player can play the file
func omxCanPlay(path string) bool {
	return RegexFormats.Match([]byte(path))
}
