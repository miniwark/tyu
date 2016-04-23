package main

import (
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/host"
)

// hostinfo represent the computer informations
type hostinfo struct {
	hostname        string // host name of the system ex 'mycomputer'
	domainname      string // domain name of the system ex 'mydomain.com'
	os              string // OS type ex 'Linux'
	osRelease       string // OS release ex '4.2.0-35-generic'
	platform        string // OS distribution or vendor ex 'Ubuntu'
	platformVersion string // OS distribution version ex '15.10'
	arch            string // architectureex. 'x86_64'
}

// Get informations about the computer by using`gopesutil` and `syscall` packages
func getHostinfo() hostinfo {
	ret := hostinfo{}

	info, err := host.Info()
	if err == nil {
		ret.hostname = info.Hostname
		ret.os = strings.Title(info.OS)
		ret.platform = strings.Title(info.Platform)
		ret.platformVersion = info.PlatformVersion
	}

	uts := syscall.Utsname{}  //TODO add an OS check because this is Linux only
	err = syscall.Uname(&uts) //TODO so we may use host.Info() for this
	if err == nil {
		ret.domainname = int8SliceToString(uts.Domainname[:])
		ret.osRelease = int8SliceToString(uts.Release[:])
		ret.arch = int8SliceToString(uts.Machine[:])
	}

	return ret
}

// Return the uptime by using `gopesutil` package in a redable string ex '10h10m01s'
func getUptime() (uptime string) {
	ret := ""
	t, err := host.Uptime()
	if err == nil {
		ret = (time.Duration(t) * time.Second).String()
	}
	return ret
}

// getUptime is in a separate func than `getHostinfo` to avoid unnecessary calls
// as all the host informations will normaly not change contrary to uptime
