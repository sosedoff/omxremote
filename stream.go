package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var durationRegexp = regexp.MustCompile(`\s?Duration: ([\d]+:[\d]+:[\d]+)`)

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

			if duration, found := parseDuration(line); found {
				s.duration = duration
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

func parseDuration(line string) (uint64, bool) {
	// It must match the output format "Duration: hh:mm:ss"
	matches := durationRegexp.FindAllStringSubmatch(line, 1)
	if len(matches) == 0 {
		log.Println("duration matches:", matches)
		return 0, false
	}

	// Read duration values
	var hours, minutes, seconds int
	_, err := fmt.Sscanf(matches[0][1], "%d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		fmt.Println("Cant parse duration line:", err)
		return 0, false
	}

	return uint64(seconds + minutes*60 + hours*3600), true
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
