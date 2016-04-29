package main

import (
	"strconv"
	"testing"

	"github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

// TestGetRaminfo test the returned fields values of `getMeminfo()`
func TestGetRaminfo(t *testing.T) {
	// setup the faking of `mem.VirtualMemory()`
	oldMemVirtualMemory := memVirtualMemory
	memVirtualMemory = func() (*mem.VirtualMemoryStat, error) {
		ret := &mem.VirtualMemoryStat{
			Total:       uint64(1024 * 1024 * 1024), // KiB to GiB conversion is implicitly tested --> ((1024 * 1024 * 1024) / (1024 * 1024 * 1024)) = 1.00
			Used:        uint64(1024 * 1024 * 1024),
			UsedPercent: float64(100),
		}
		return ret, nil
	}

	// test
	expected := raminfo{
		total:       strconv.FormatFloat(1.00, 'f', 2, 64),
		used:        strconv.FormatFloat(1.00, 'f', 2, 64),
		usedPercent: int(100),
	}
	actual, err := getRaminfo()

	assert.NoError(t, err, "`getRaminfo()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getRaminfo` should be equal to main.raminfo{total:\"1.00\", used:\"1.00\", usedPercent:100}")

	// teardown
	memVirtualMemory = oldMemVirtualMemory
}

// TestGetRaminfoType test if `getRaminfo()` return a `raminfo` type and if each fields have the correct types
func TestGetRaminfoType(t *testing.T) {
	expected := raminfo{
		total:       "", // the result values of the fields are not tested
		used:        "",
		usedPercent: int(0),
	}
	actual, _ := getRaminfo()

	assert.IsType(t, expected, actual, "`getRaminfo()` should return a `raminfo` type")
	assert.IsType(t, expected.total, actual.total, "`getRaminfo()` should return a `total` field with a string type")
	assert.IsType(t, expected.used, actual.used, "`getRaminfo()` should return a `used` field with a string type")
	assert.IsType(t, expected.usedPercent, actual.usedPercent, "`getRaminfo()` should return a `usedPercent` field with an int type")
}

// TestGetSwapinfo test the returned fields values of `getSwapinfo()`
func TestGetSwapinfo(t *testing.T) {
	// setup the faking of `mem.SwapMemory()`
	OldMemSwapMemory := memSwapMemory
	memSwapMemory = func() (*mem.SwapMemoryStat, error) {
		ret := &mem.SwapMemoryStat{
			Total:       uint64(1024 * 1024 * 1024),
			Used:        uint64(1024 * 1024 * 1024),
			UsedPercent: float64(100),
		}
		return ret, nil
	}

	// test
	expected := swapinfo{
		total:       strconv.FormatFloat(1.00, 'f', 2, 64),
		used:        strconv.FormatFloat(1.00, 'f', 2, 64),
		usedPercent: int(100),
	}
	actual, err := getSwapinfo()

	assert.NoError(t, err, "`getSwapinfo()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getSwapinfo()` should be equal to main.swapinfo{total:\"1.00\", used:\"1.00\", usedPercent:100}")

	// teardown
	memSwapMemory = OldMemSwapMemory
}

// TestGetSwapinfoType test if `getSwapinfo()` return a `swapinfo` type and if each fields have the correct types
func TestGetSwapinfoType(t *testing.T) {
	expected := swapinfo{
		total:       "", // the result values of the fields are not tested
		used:        "",
		usedPercent: int(0),
	}
	actual, _ := getSwapinfo()

	assert.IsType(t, expected, actual, "`getSwapinfo()` should return a `swapinfo` type")
	assert.IsType(t, expected.total, actual.total, "`getSwapinfo()` should return a `ramTotal` field with a string type")
	assert.IsType(t, expected.used, actual.used, "`getSwapinfo()` should return a `ramUsed` field with a string type")
	assert.IsType(t, expected.usedPercent, actual.usedPercent, "`getSwapinfo()` should return a `ramUsedPercent` field with an int type")
}

// TestMemVirtualMemory test if `memVirtualMemory()` return a value with a *mem.VirtualMemoryStat type
func TestMemVirtualMemory(t *testing.T) {
	expected := &mem.VirtualMemoryStat{}
	actual, _ := memVirtualMemory()

	assert.IsType(t, expected, actual, "`memVirtualMemory()` should return a *mem.VirtualMemoryStat")
}

// TestMemSwapMemory test if `memSwapMemory()` return a value with a *mem.SwapMemoryStat type
func TestMemSwapMemory(t *testing.T) {
	expected := &mem.SwapMemoryStat{}
	actual, _ := memSwapMemory()

	assert.IsType(t, expected, actual, "`memSwapMemory()` should return a *mem.SwapMemoryStat type")
}

//TODO add tests for errors --> must return empty/zero values
