package main

import (
	"fmt"
	"regexp"
)

var (
	storageRe = regexp.MustCompile(`\/dev\/root\s+([\d]+)\s+([\d]+)\s+([\d]+)\s+([\d]+%)`)
	memoryRe  = regexp.MustCompile(`Mem:\s+([\d]+)\s+([\d]+)\s+([\d]+)`)
)

type HostInfo struct {
	Parsed      bool `json:"parsed"`
	Total       int  `json:"total"`
	Used        int  `json:"used"`
	Available   int  `json:"available"`
	UsedPercent int  `json:"used_percent"`
}

func parseStorageInfo(input string) HostInfo {
	info := HostInfo{}
	result := storageRe.FindAllStringSubmatch(input, -1)

	if len(result) == 1 {
		info.Parsed = true
		fmt.Sscanf(result[0][1], "%d", &info.Total)
		fmt.Sscanf(result[0][2], "%d", &info.Used)
		fmt.Sscanf(result[0][3], "%d", &info.Available)
		fmt.Sscanf(result[0][4], "%d", &info.UsedPercent)
	}

	return info
}

func parseMemoryInfo(input string) HostInfo {
	info := HostInfo{}
	result := memoryRe.FindAllStringSubmatch(input, -1)

	if len(result) == 1 {
		info.Parsed = true

		fmt.Sscanf(result[0][1], "%d", &info.Total)
		fmt.Sscanf(result[0][2], "%d", &info.Used)
		fmt.Sscanf(result[0][3], "%d", &info.Available)

		if info.Total > 0 {
			info.UsedPercent = (info.Used * 100) / info.Total
		}
	}

	return info
}

func parseTemperature(input string) float32 {
	var value float32
	fmt.Sscanf(input, "%v", &value)
	return value
}
