package main

import "github.com/shirou/gopsutil/net"

// Netinfo represent the network traffic informations
type Netinfo struct {
	up   float64
	down float64
}

// Get informations about the net traffic
func getNetinfo() Netinfo {
	ret := Netinfo{}

	iocounters, err := net.IOCounters(false)
	if err == nil {
		ret.up = float64(iocounters[0].BytesSent) / 1024
		ret.down = float64(iocounters[0].BytesRecv) / 1024
		//ret.up = strconv.FormatFloat(float64(ioconters[0].BytesSent)/1024, 'f', 2, 64),
		//ret.down = strconv.FormatFloat(float64(ioconters[0].BytesRecv)/1024, 'f', 2, 64),
	}
	return ret
}

//TODO try to move the (networkNew - networkOld) calculations from main.go here
// and put back Netinfo as netinfo with strings vars
