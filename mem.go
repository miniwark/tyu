package main

import (
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

// meminfo represent the memory usage statistics
type meminfo struct {
	total       string // total available memory in gigabytes
	used        string // used memory in gigabytes
	usedPercent int    // used memory in percents of total memory

}

// Get RAM usage informations by using `gopesutil` package
// and then convert them to `string` or `int`
func getRaminfo() (ret meminfo, err error) {
	ram, err := memVirtualMemory()
	if err == nil {
		ret.total = strconv.FormatFloat(float64(ram.Total)/(1024*1024*1024), 'f', 2, 64) // (1024*1024*1024) to convert to GiB from `gopesutil`
		ret.used = strconv.FormatFloat(float64(ram.Used)/(1024*1024*1024), 'f', 2, 64)
		ret.usedPercent = int(ram.UsedPercent)
	}

	return ret, err
}

// Get Swap usage informations by using `gopesutil` package
// and then convert them to `string` or `int`
func getSwapinfo() (ret meminfo, err error) {
	swap, err := memSwapMemory()
	if err == nil {
		ret.total = strconv.FormatFloat(float64(swap.Total)/(1024*1024*1024), 'f', 2, 64)
		ret.used = strconv.FormatFloat(float64(swap.Used)/(1024*1024*1024), 'f', 2, 64)
		ret.usedPercent = int(swap.UsedPercent)
	}

	return ret, err
}

// wrap `mem.VirtualMemory()` in an unexported variable for testability
var memVirtualMemory = func() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// wrap `mem.SwapMemory()` in an unexported variable for testability
var memSwapMemory = func() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
