package main

import (
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/host"
)

// `hostinfo` represent the computer informations
type hostinfo struct {
	hostname        string // host name of the system ex 'mycomputer'
	domainname      string // domain name of the system ex 'mydomain.com'
	os              string // OS type ex 'Linux'
	osRelease       string // OS release ex '4.2.0-35-generic'
	platform        string // OS distribution or vendor ex 'Ubuntu'
	platformVersion string // OS distribution version ex '15.10'
	arch            string // architectureex. 'x86_64'
}

// Get informations about the computer by using `syscall`and `gopesutil` packages
func getHostinfo() hostinfo {

	uts := syscall.Utsname{}   //TODO add an OS check because this is not portable
	err := syscall.Uname(&uts) //TODO so we may use host.Info() for this
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	info, err := host.Info()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	result := hostinfo{
		hostname:        info.Hostname,
		domainname:      int8SliceToString(uts.Domainname[:]),
		os:              strings.Title(info.OS),
		osRelease:       int8SliceToString(uts.Release[:]),
		platform:        strings.Title(info.Platform),
		platformVersion: info.PlatformVersion,
		arch:            int8SliceToString(uts.Machine[:]),
	}
	return result
}

// Return the uptime by using `gopesutil` package in a redable string ex '10h10m01s'
func getUptime() (uptime string) {
	t, err := host.Uptime()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	uptime = (time.Duration(t) * time.Second).String()
	return uptime
}

// getUptime is in a separate func than `getHostinfo` to avoid unecessary calls
// as all the host informations will normaly not change contrary to uptime
