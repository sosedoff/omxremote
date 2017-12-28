package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

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
		OmxPath,       // path to omxplayer executable
		"--stats",     // print stats to stdout (buffers, time, etc)
		"--with-info", // print stats about streams before playback
		"--refresh",   // adjust framerate/resolution to video
		"--blank",     // set background to black
		"--adev",      // audio out device
		"hdmi",        // using hdmi for audio/video
		file,          // path to video file
	)

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
	if RegexFormats.Match([]byte(path)) {
		return true
	}

	return false
}
