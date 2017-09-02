package main

import (
	"log"

	"github.com/grandcat/zeroconf"
)

const (
	zeroConfName    = "app"
	zeroconfService = "_omxremote._tcp"
	zeroconfDomain  = "local."
	zeroconfPort    = 8080
)

func startZeroConfAdvertisement(stop chan bool) {
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
