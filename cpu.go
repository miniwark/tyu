package main

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// cpuinfo represent the CPU informations
type cpuinfo struct {
	count     string // number of CPUs
	vendorID  string // vendor name ex. AuthenticAMD, GenuineIntel
	modelName string // CPU model name
	cpuMhz    string // frequency of the CPU in Mhz
	//TODO temperature: see lm-sensors
}

// Get informations about the cpu by using `gopesutil` package
func getCPUinfo() cpuinfo {
	ret := cpuinfo{}

	info, err := cpu.Info() // cpu.Info() return a slice of InfoStat structs
	if err == nil {
		ret.count = strconv.Itoa(len(info))
		ret.vendorID = info[0].VendorID
		ret.modelName = info[0].ModelName
		ret.cpuMhz = strconv.FormatFloat(info[0].Mhz, 'f', 0, 64)
	}
	return ret
}

// get the system-wide CPU utilization percentage
func getCPUpercent() int {
	percent, err := cpu.Percent((500 * time.Millisecond), false) // 0.5 seconds, `false` for system wide
	if err == nil {
		return int(percent[0]) // even if cpu.Percent() use `false` it return a slice
	}
	return 0
}

// `getCPUpercent` is in a separate func than `getCPUinfo` to avoid unnecessary calls
// as all the host informations will normaly not change contrary to uptime

//TODO count cores ?
//TODO count the physicals id = multiples sockets

//TODO We need to improve this. For now it take vendorID, modelName, etc from
//the first CPU assuming than all CPUs are the same (probably OK)and than there
//is only one socket (not true on servers)
//
// http://superuser.com/questions/388115/interpreting-output-of-cat-proc-cpuinfo
// http://unix.stackexchange.com/questions/146051/number-of-processors-in-proc-cpuinfo

// Get informations about the cpu by using `gopesutil` package
