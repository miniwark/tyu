package main

import (
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

// `meminfo` represent the memory usage statistics
type meminfo struct {
	memTotal        string // total available memory in megabytes
	memUsed         string // used memory in megabytes
	memFree         string // free memory in megabytes
	memUsedPercent  int    // used memory in percents of total memory
	swapTotal       string // total available swap memory in megabytes
	swapUsed        string // used swap memory in megabytes
	swapFree        string // free swap memory in megabytes
	swapUsedPercent int    // used swap memory in percents of total memory
}

// Get memory usage informations by using `gopesutil` package
// and then convert them to `string` or `int`
func getMeminfo() meminfo {
	virtual, err := mem.VirtualMemory()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	swap, err := mem.SwapMemory()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	result := meminfo{
		memTotal:        strconv.FormatUint(virtual.Total, 10),
		memUsed:         strconv.FormatUint(virtual.Used, 10),
		memFree:         strconv.FormatUint(virtual.Free, 10),
		memUsedPercent:  int(virtual.UsedPercent),
		swapTotal:       strconv.FormatUint(swap.Total, 10),
		swapUsed:        strconv.FormatUint(swap.Total, 10),
		swapFree:        strconv.FormatUint(swap.Total, 10),
		swapUsedPercent: int(swap.UsedPercent),
	}
	return result
}
