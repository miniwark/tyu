package main

import (
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

// meminfo represent the memory usage statistics
type meminfo struct {
	ramTotal        string // total available memory in gigabytes
	ramUsed         string // used memory in gigabytes
	ramUsedPercent  int    // used memory in percents of total memory
	swapTotal       string // total available swap memory in gigabytes
	swapUsed        string // used swap memory in gigabytes
	swapUsedPercent int    // used swap memory in percents of total memory
}

// Get memory usage informations by using `gopesutil` package
// and then convert them to `string` or `int`
func getMeminfo() meminfo {
	ret := meminfo{}

	ram, err := mem.VirtualMemory()
	if err == nil {
		ret.ramTotal = strconv.FormatFloat(float64(ram.Total)/(1024*1024*1024), 'f', 2, 64) // (1024*1024*1024) to convert to GiB from `gopesutil`
		ret.ramUsed = strconv.FormatFloat(float64(ram.Used)/(1024*1024*1024), 'f', 2, 64)
		ret.ramUsedPercent = int(ram.UsedPercent)
	} else {
		ret.ramTotal = ""
		ret.ramUsed = ""
		ret.ramUsedPercent = 0
	}

	swap, err := mem.SwapMemory()
	if err == nil {
		ret.swapTotal = strconv.FormatFloat(float64(swap.Total)/(1024*1024*1024), 'f', 2, 64)
		ret.swapUsed = strconv.FormatFloat(float64(swap.Used)/(1024*1024*1024), 'f', 2, 64)
		ret.swapUsedPercent = int(swap.UsedPercent)
	} else {
		ret.swapTotal = ""
		ret.swapUsed = ""
		ret.swapUsedPercent = 0
	}

	return ret
}
