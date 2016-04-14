package main

import (
	"strconv"

	"github.com/shirou/gopsutil/net"
)

// `netinfo` represent the network trafic informations
type netinfo struct {
	up   string
	down string
}

// Get informations about the net trafic
func getNetinfo() netinfo {
	ioconters, err := net.IOCounters(false)
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	ret := netinfo{
		up:   strconv.FormatUint(ioconters[1].BytesSent, 10),
		down: strconv.FormatUint(ioconters[1].BytesRecv, 10),
	}
	return ret
}
