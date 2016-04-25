package main

import (
	"strconv"
	"testing"

	"github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

// TestGetMeminfo test the returned fields values of `getMeminfo()`
func TestGetMeminfo(t *testing.T) {
	// setup the faking of `mem.VirtualMemory()` & `mem.SwapMemory()`
	oldMemVirtualMemory := memVirtualMemory
	OldMemSwapMemory := memSwapMemory
	memVirtualMemory = func() (*mem.VirtualMemoryStat, error) {
		ret := &mem.VirtualMemoryStat{
			Total:       uint64(1024 * 1024 * 1024), // KiB to GiB conversion is implicitly tested --> ((1024 * 1024 * 1024) / (1024 * 1024 * 1024)) = 1.00
			Used:        uint64(1024 * 1024 * 1024),
			UsedPercent: float64(100),
		}
		return ret, nil
	}
	memSwapMemory = func() (*mem.SwapMemoryStat, error) {
		ret := &mem.SwapMemoryStat{
			Total:       uint64(1024 * 1024 * 1024),
			Used:        uint64(1024 * 1024 * 1024),
			UsedPercent: float64(100),
		}
		return ret, nil
	}

	// test
	expected := meminfo{
		ramTotal:        strconv.FormatFloat(1.00, 'f', 2, 64),
		ramUsed:         strconv.FormatFloat(1.00, 'f', 2, 64),
		ramUsedPercent:  int(100),
		swapTotal:       strconv.FormatFloat(1.00, 'f', 2, 64),
		swapUsed:        strconv.FormatFloat(1.00, 'f', 2, 64),
		swapUsedPercent: int(100),
	}
	actual := getMeminfo()

	assert.Equal(t, expected, actual, "`getMeminfo` should be equal to main.Netinfo{ramTotal:\"1.00\", ramUsed:\"100\", ramUsedPercent:100, swapTotal:\"1.00\", swapUsed:\"1.00\", swapUsedPercent:100}")

	// teardown
	memVirtualMemory = oldMemVirtualMemory
	memSwapMemory = OldMemSwapMemory
}

// TestGetMeminfoType test if `getMeminfo()` return a `meminfo` type and if each fields have the correct types
func TestGetMeminfoType(t *testing.T) {
	expected := meminfo{
		ramTotal:        "", // the result values of the fields are not tested
		ramUsed:         "",
		ramUsedPercent:  int(0),
		swapTotal:       "",
		swapUsed:        "",
		swapUsedPercent: int(0),
	}
	actual := getMeminfo()

	assert.IsType(t, expected, actual, "`getMeminfo()` should return a `meminfo` type")
	assert.IsType(t, expected.ramTotal, actual.ramTotal, "`getMeminfo()` should return a `ramTotal` field with a string type")
	assert.IsType(t, expected.ramUsed, actual.ramUsed, "`getMeminfo()` should return a `ramUsed` field with a string type")
	assert.IsType(t, expected.ramUsedPercent, actual.ramUsedPercent, "`getMeminfo()` should return a `ramUsedPercent` field with an int type")
	assert.IsType(t, expected.swapTotal, actual.swapTotal, "`getMeminfo()` should return a `swapTotal` field with a string type")
	assert.IsType(t, expected.swapUsed, actual.swapUsed, "`getMeminfo()` should return a `swapUsed` field with a string type")
	assert.IsType(t, expected.swapUsedPercent, actual.swapUsedPercent, "`getMeminfo()` should return a `swapUsedPercent` field with an int type")
}

//TODO add tests for errors --> must return empty/zero values
