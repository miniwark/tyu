package main

import (
	"time"

	"github.com/shirou/gopsutil/host"
)

//TODO add bootime with a struct ?

// Return the uptime by using `gopesutil` package in a redable string ex '10h10m01s'
func getUptime() (uptime string) {
	t, err := host.Uptime()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	uptime = (time.Duration(t) * time.Second).String()
	return uptime
}
