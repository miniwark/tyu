package main

import (
	"strconv"

	"github.com/shirou/gopsutil/net"
)

// Netinfo represent the network trafic informations
type Netinfo struct {
	up   string
	down string
}

// Get informations about the net trafic
func getNetinfo() Netinfo {
	ioconters, err := net.IOCounters(false)
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	ret := Netinfo{
		//up:   strconv.FormatUint(ioconters[0].BytesSent, 10),
		up: strconv.FormatFloat(float64(ioconters[0].BytesSent)/(1024*1024), 'f', 2, 64),
		//down: strconv.FormatUint(ioconters[0].BytesRecv, 10),
		down: strconv.FormatFloat(float64(ioconters[0].BytesRecv)/(1024*1024), 'f', 2, 64),
	}
	return ret
}
