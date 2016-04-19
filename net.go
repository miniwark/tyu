package main

import "github.com/shirou/gopsutil/net"

// Netinfo represent the network trafic informations
type Netinfo struct {
	up   float64
	down float64
}

// Get informations about the net trafic
func getNetinfo() Netinfo {
	ioconters, err := net.IOCounters(false)
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	ret := Netinfo{
		//up: strconv.FormatFloat(float64(ioconters[0].BytesSent)/(1024*1024), 'f', 2, 64),
		//up: float64(ioconters[0].BytesSent) / (1024 * 1024),
		up: float64(ioconters[0].BytesSent) / 1024,
		//down: strconv.FormatFloat(float64(ioconters[0].BytesRecv)/(1024*1024), 'f', 2, 64),
		//down: float64(ioconters[0].BytesRecv) / (1024 * 1024),
		down: float64(ioconters[0].BytesRecv) / 1024,
	}
	return ret
}

//TODO try to move the (networkNew - networkOld) calculations from main.go here
// and put back Netinfo as netinfo with strings vars
