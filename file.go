package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Returns true if the specified file exists
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

// Convert media filename to regular title
func fileToTitle(name string) string {
	// Remove file extension from name
	name = strings.Replace(name, filepath.Ext(name), "", 1)

	// Replace all dots with white space
	name = strings.Replace(name, ".", " ", -1)

	// Replace parenteses and brackets
	name = RegexBrackets.ReplaceAllString(name, "")

	// Check if name has a typical torrent format: "name year format etc"
	pos := RegexYear.FindStringIndex(name)
	if len(pos) > 0 {
		name = name[0:pos[0]]
	} else {
		// Check if its an episode, i.e. "S##E##"
		pos = RegexEpisode.FindStringIndex(name)

		if len(pos) > 0 {
			name = name[0:pos[0]]
		}
	}

	// Remove junk stuff
	name = RegexJunk.ReplaceAllString(name, "")

	// Remove any extra white space
	name = RegexSpace.ReplaceAllString(name, "")

	// Truncate space
	name = strings.TrimSpace(name)

	return name
}
