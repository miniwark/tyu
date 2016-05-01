package main

import (
	"strconv"
	"testing"

	"github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

// TestGetRAMStat test the returned fields values of `getRAMStat()`
func TestGetRAMStat(t *testing.T) {
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
	expected := memStat{
		total:       strconv.FormatFloat(1.00, 'f', 2, 64),
		used:        strconv.FormatFloat(1.00, 'f', 2, 64),
		usedPercent: int(100),
	}
	actual, err := getRAMStat()

	assert.NoError(t, err, "`getRAMStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getRAMStat` should be equal to main.memStat{total:\"1.00\", used:\"1.00\", usedPercent:100}")

	// teardown
	memVirtualMemory = oldMemVirtualMemory
}

// TestGetRAMStatType test if `getRAMStat()` return a `memStat` type and if each fields have the correct types
func TestGetRAMStatType(t *testing.T) {
	expected := memStat{
		total:       "", // the result values of the fields are not tested
		used:        "",
		usedPercent: int(0),
	}
	actual, _ := getRAMStat()

	assert.IsType(t, expected, actual, "`getRAMStat()` should return a `memStat` type")
	assert.IsType(t, expected.total, actual.total, "`getRAMStat()` should return a `total` field with a string type")
	assert.IsType(t, expected.used, actual.used, "`getRAMStat()` should return a `used` field with a string type")
	assert.IsType(t, expected.usedPercent, actual.usedPercent, "`getRAMStat()` should return a `usedPercent` field with an int type")
}

// TestGetSwapStat test the returned fields values of `getSwapStat()`
func TestGetSwapStat(t *testing.T) {
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
	expected := memStat{
		total:       strconv.FormatFloat(1.00, 'f', 2, 64),
		used:        strconv.FormatFloat(1.00, 'f', 2, 64),
		usedPercent: int(100),
	}
	actual, err := getSwapStat()

	assert.NoError(t, err, "`getSwapStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getSwapStat()` should be equal to main.memStat{total:\"1.00\", used:\"1.00\", usedPercent:100}")

	// teardown
	memSwapMemory = OldMemSwapMemory
}

// TestGetSwapStatType test if `getSwapStat()` return a `memStat` type and if each fields have the correct types
func TestGetSwapStatType(t *testing.T) {
	expected := memStat{
		total:       "", // the result values of the fields are not tested
		used:        "",
		usedPercent: int(0),
	}
	actual, _ := getSwapStat()

	assert.IsType(t, expected, actual, "`getSwapStat()` should return a `memStat` type")
	assert.IsType(t, expected.total, actual.total, "`getSwapStat()` should return a `total` field with a string type")
	assert.IsType(t, expected.used, actual.used, "`getSwapStat()` should return a `used` field with a string type")
	assert.IsType(t, expected.usedPercent, actual.usedPercent, "`getSwapStat()` should return a `usedPercent` field with an int type")
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
