package main

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// cpuStat represent the CPU informations
type cpuStat struct {
	count     string // number of CPUs
	vendorID  string // vendor name ex. AuthenticAMD, GenuineIntel
	modelName string // CPU model name
	mhz       string // frequency of the CPU in Mhz
}

//TODO var temperature: see lm-sensors

// getCPUStat get informations about the cpu by using `gopesutil` package
func getCPUStat() (ret cpuStat, err error) {
	cpu, err := cpuInfo() // cpu.Info() return a slice of InfoStat structs
	if err == nil {
		ret.count = strconv.Itoa(len(cpu))
		ret.vendorID = cpu[0].VendorID
		ret.modelName = cpu[0].ModelName
		ret.mhz = strconv.FormatFloat(cpu[0].Mhz, 'f', 0, 64)
	}
	return ret, err
}

// getCPUPercent get the system-wide CPU utilization percentage
func getCPUPercent() (ret int, err error) {
	percent, err := cpuPercent((500 * time.Millisecond), false) // 0.5 seconds, `false` for system wide
	if err == nil {
		ret = int(percent[0]) // even if cpu.Percent() use `false` it return a slice
	}
	return ret, err
}

// `getCPUPercent` is in a separate func than `getCPUinfo` to avoid unnecessary calls
// as all the host informations will normaly not change contrary to CPU percent usage

// wrap `cpu.Info()` in an unexported variable for testability
var cpuInfo = func() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// wrap `cpu.Percent()` in an unexported variable for testability
var cpuPercent = func(interval time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(interval, percpu)
}

//TODO count cores ?
//TODO count the physicals id = multiples sockets

//TODO We need to improve this. For now it take vendorID, modelName, etc from
//the first CPU assuming than all CPUs are the same (probably OK)and than there
//is only one socket (not true on servers)
//
// http://superuser.com/questions/388115/interpreting-output-of-cat-proc-cpuinfo
// http://unix.stackexchange.com/questions/146051/number-of-processors-in-proc-cpuinfo

// Get informations about the cpu by using `gopesutil` package
