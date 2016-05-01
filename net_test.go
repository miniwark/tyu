package main

import (
	"testing"

	"github.com/shirou/gopsutil/net"
	"github.com/stretchr/testify/assert"
)

// TestGetNetStat test the returned fields values of `getNetStat()`
func TestGetNetStat(t *testing.T) {
	// setup the faking of `net.IOCounters()`
	oldNetIocounters := netIocounters
	netIocounters = func(pernic bool) ([]net.IOCountersStat, error) {
		ret := []net.IOCountersStat{
			{
				BytesSent: uint64(1024), // Bytes to KiB conversion is implicitly tested --> (1024 / 1024) = 1
				BytesRecv: uint64(1024),
			},
		}
		return ret, nil
	}

	// test
	expected := NetStat{
		up:   float64(1),
		down: float64(1),
	}
	actual, err := getNetStat()

	assert.NoError(t, err, "`getNetStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getNetStat()` should be equal to main.NetStat{up:1, down:1}")

	// teardown
	netIocounters = oldNetIocounters
}

// TestGetNetStatType test if `getNetStat()` return a `NetStat` type and if each fields have the correct types
func TestGetNetStatType(t *testing.T) {
	expected := NetStat{
		up:   float64(0), // the result values of the fields are not tested
		down: float64(0),
	}
	actual, _ := getNetStat()

	assert.IsType(t, expected, actual, "`getNetStat()` should return a `NetStat` type")
	assert.IsType(t, expected.up, actual.up, "`getNetStat()` should return a `up` field with a float64 type")
	assert.IsType(t, expected.down, actual.down, "`getNetStat()` should return a `down` field with a float64 type")
}

// TestNetIocounters test if `netIocounters()` return a value with a []net.IOCountersStat slice
func TestNetIocounters(t *testing.T) {
	expected := []net.IOCountersStat{}
	actual, _ := netIocounters(false)

	assert.IsType(t, expected, actual, "`netIocounters()` should return a []net.IOCountersStat slice")
}

//TODO add tests for errors --> must return empty/zero values
