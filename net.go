package main

import (
	"strconv"

	"github.com/shirou/gopsutil/net"
)

// `netinfo` represent the network trafic informations
type netinfo struct {
	out string
	in  string
}

// Get informations about the net trafic
func getNetinfo() netinfo {
	ioconters, err := net.IOCounters(false)
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	ret := netinfo{
		out: strconv.FormatUint(ioconters[0].BytesSent, 10),
		in:  strconv.FormatUint(ioconters[0].BytesRecv, 10),
	}
	return ret
}
