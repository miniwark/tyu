package main

import "github.com/shirou/gopsutil/host"

// `hostinfo` represent the computer informations
type hostinfo struct { //TODO rename and avoid `sysinfo`
	hostname string // host name of the system ex 'mycomputer'
	//domainname string // domain name of the system ex 'mydomain.com' //TODO
	os string // OS type ex 'Linux'
	//osRelease  string // OS release ex '4.2.0-35-generic' //TODO
	platform        string // OS distribution or vendor ex 'Ubuntu'
	platformVersion string // OS distribution version ex '15.10'
}

// Get informations about the computer by using `gopesutil` package
func getHostinfo() hostinfo {
	info, err := host.Info()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	result := hostinfo{
		hostname: info.Hostname,
		//domainame //TODO using `syscall` ?
		os: info.OS,
		//osRelease
		platform:        info.Platform,
		platformVersion: info.PlatformVersion,
	}
	return result
}
