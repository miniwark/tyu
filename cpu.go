package main

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// `cpuinfo` represent the CPU informations
type cpuinfo struct {
	count     string // number of CPUs
	vendorID  string // vendor name ex. AuthenticAMD, GenuineIntel
	modelName string // CPU model name
	cpuMhz    string // speed of the CPU
	//TODO temperature and fan speed ?
}

// Get informations about the cpu by using `gopesutil` packages
func getCPUinfo() cpuinfo {
	info, err := cpu.Info() // cpu.Info() return a slice of InfoStat structs
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	result := cpuinfo{
		count:     strconv.Itoa(len(info)),
		vendorID:  info[0].VendorID, //BUG in `gopsutil`
		modelName: info[0].ModelName,
		cpuMhz:    strconv.FormatFloat(info[0].Mhz, 'f', 0, 64),
	}
	return result
}

// get the system-wide CPU utilization percentage
func getCPUpercent() (usedPercent int) {
	percent, err := cpu.Percent((500 * time.Millisecond), false) // 0.5 seconds, `false` for system wide
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	return int(percent[0]) // even if cpu.Percent() use `false` it return a slice
}

// `getCPUpercent` is in a separate func than `getCPUinfo` to avoid unecessary calls
// as all the host informations will normaly not change contrary to uptime

//TODO count cores ?
//TODO count the physicals id = multiples sockets

//TODO We need to improve this. For now it take vendorID, modelName, etc from
//the first CPU assuming than all CPUs are the same (probably OK)and than there
//is only one socket (not true on servers)
//
// http://superuser.com/questions/388115/interpreting-output-of-cat-proc-cpuinfo
// http://unix.stackexchange.com/questions/146051/number-of-processors-in-proc-cpuinfo
//
// TODO add temperature
