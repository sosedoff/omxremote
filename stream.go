package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var durationRegexp = regexp.MustCompile(`\s?Duration: [\d]+:[\d]+:[\d]+`)

func durationFromSeconds(value uint64) string {
	hours := value / 3600
	minutes := (value - hours*3600) / 60
	seconds := value - (hours*3600 + minutes*60)

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

type Position struct {
	nanos   uint64
	seconds uint64
}

func (p *Position) Set(value uint64) {
	p.nanos = value
	p.seconds = value / 1000000
}

func (p *Position) String() string {
	return durationFromSeconds(p.seconds)
}

type Stream struct {
	duration uint64
	pos      Position
}

func NewStream() *Stream {
	return &Stream{
		pos:      Position{},
		duration: 0,
	}
}

func (s *Stream) Start(stdout, stderr io.Reader) {
	s.duration = 0
	s.pos.Set(0)

	// Most meta information comes from STDERR
	go func() {
		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			line := scanner.Text()
			if s.parseDuration(line) {
				break
			}
		}
	}()

	progress := bufio.NewReader(stdout)

	for {
		data, err := progress.ReadBytes('\r')
		if err != nil {
			log.Println("ERROR:", err)
			break
		}
		s.parsePosition(string(data))
	}
}

func (s *Stream) parseDuration(line string) bool {
	// It must match the output format "Duration: hh:mm:ss"
	if !durationRegexp.MatchString(line) {
		return false
	}

	// There's a bunch of junk at the beginning of the string
	line = strings.TrimSpace(line)

	// Read duration values
	var hours, minutes, seconds int
	_, err := fmt.Sscanf(line, "Duration: %d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		fmt.Println("Cant parse duration line:", err)
		return true
	}

	s.duration = uint64(seconds + minutes*60 + hours*3600)
	return true
}

func (s *Stream) parsePosition(line string) {
	if !strings.HasPrefix(line, "M:") {
		return
	}

	var posNanos uint64
	if n, _ := fmt.Sscanf(line, "M:%d", &posNanos); n == 1 {
		s.pos.Set(posNanos)
	}
}
