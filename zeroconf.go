package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grandcat/zeroconf"
)

var (
	zeroConfName    = "Omxremote"
	zeroconfService = "_omxremote._tcp"
	zeroconfDomain  = "local."
	zeroconfPort    = 8080
)

func startZeroConfAdvertisement(stop chan bool) {
	hostname, err := os.Hostname()
	if err == nil && hostname != "" {
		zeroConfName = fmt.Sprintf("%s (%s)", zeroConfName, strings.Split(hostname, ".")[0])
	}

	log.Println("Starting zeroconf:", zeroConfName, zeroconfService, zeroconfPort)
	defer log.Println("Zeroconf service terminated")

	server, err := zeroconf.Register(
		zeroConfName,
		zeroconfService,
		zeroconfDomain,
		zeroconfPort,
		[]string{"version=" + VERSION},
		nil,
	)
	if err != nil {
		log.Println("Zeroconf server error:", err)
		return
	}
	defer server.Shutdown()

	<-stop
}
