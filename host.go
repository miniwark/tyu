package main

import (
	"strings"
	"syscall" //TODO replace by  golang.org/x/sys
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
func getHostinfo() (ret hostinfo, err error) {
	info, err := hostInfo()
	if err == nil {
		ret.hostname = info.Hostname
		ret.os = strings.Title(info.OS)
		ret.platform = strings.Title(info.Platform)
		ret.platformVersion = info.PlatformVersion
	}

	if ret.os == "Linux" {
		uts, err1 := getUname() //TODO move this un a separate func?
		if err1 == nil {
			ret.domainname = int8SliceToString(uts.Domainname[:])
			ret.osRelease = int8SliceToString(uts.Release[:])
			ret.arch = int8SliceToString(uts.Machine[:])
		} else {
			err = appendError(err, err1)
		}
	}

	return ret, err

	//TODO add uname like information from other OS
}

// Return the uptime by using `gopesutil` package in a redable string ex '10h10m01s'
func getUptime() (ret string, err error) {
	ret = ""
	t, err := hostUptime()
	if err == nil {
		ret = (time.Duration(t) * time.Second).String()
	}
	return ret, err
}

// getUptime is in a separate func than `getHostinfo` to avoid unnecessary calls
// as all the host informations will normaly not change contrary to uptime

// Get informations from syscall Uname()
// wrapped in an an unexported variable for testability
var getUname = func() (syscall.Utsname, error) { //TODO tested with Linux only
	uts := syscall.Utsname{}
	err := syscall.Uname(&uts)
	return uts, err
}

// wrap `host.Info()` in an unexported variable for testability
var hostInfo = func() (*host.InfoStat, error) { //TODO rename
	return host.Info()
}

// wrap `host.Uptime()` in an unexported variable for testability
var hostUptime = func() (uint64, error) {
	return host.Uptime()
}
