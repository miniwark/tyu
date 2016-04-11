package main

import (
	"syscall"

	"github.com/shirou/gopsutil/host"
)

// `hostinfo` represent the computer informations
type hostinfo struct { //TODO rename and avoid `sysinfo`
	hostname        string // host name of the system ex 'mycomputer'
	domainname      string // domain name of the system ex 'mydomain.com' //TODO
	os              string // OS type ex 'Linux'
	osRelease       string // OS release ex '4.2.0-35-generic' //TODO
	platform        string // OS distribution or vendor ex 'Ubuntu'
	platformVersion string // OS distribution version ex '15.10'
	arch            string // architectureex. 'x86_64' //TODO
}

// Get informations about the computer by using `syscall`and `gopesutil` packages
func getHostinfo() hostinfo {

	uts := syscall.Utsname{} //TODO add os check because this is not portable
	err := syscall.Uname(&uts)
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	info, err := host.Info()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	result := hostinfo{
		hostname:        info.Hostname,
		domainname:      int8ArrayToString(uts.Domainname[:]),
		os:              info.OS,
		osRelease:       int8ArrayToString(uts.Release[:]),
		platform:        info.Platform,
		platformVersion: info.PlatformVersion,
		arch:            int8ArrayToString(uts.Machine[:]),
	}
	return result
}

// uptime is in a separate func than `getHostinfo` to avoid unnecessary calls
// as all the host informations will normaly not change contrary to uptime
