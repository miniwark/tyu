package main

import (
	"strconv"

	"github.com/shirou/gopsutil/cpu"
)

// `cpuinfo` represent the CPU informations
type cpuinfo struct {
	count     string // number of CPUs
	vendorID  string // vendor name ex. AuthenticAMD, GenuineIntel
	modelName string // CPU model name
	cpuMhz    string //
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
