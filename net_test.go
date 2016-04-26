package main

import (
	"testing"

	"github.com/shirou/gopsutil/net"
	"github.com/stretchr/testify/assert"
)

// TestGetNetinfo test the returned fields values of `getNetinfo()`
func TestGetNetinfo(t *testing.T) {
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
	expected := Netinfo{
		up:   float64(1),
		down: float64(1),
	}
	actual := getNetinfo()

	assert.Equal(t, expected, actual, "`getNetinfo()` should be equal to main.Netinfo{up:1, down:1}")

	// teardown
	netIocounters = oldNetIocounters
}

// TestGetNetinfoType test if `getNetinfo()` return a `Netinfo` type and if each fields have the correct types
func TestGetNetinfoType(t *testing.T) {
	expected := Netinfo{
		up:   float64(0), // the result values of the fields are not tested
		down: float64(0),
	}
	actual := getNetinfo()

	assert.IsType(t, expected, actual, "`getNetinfo()` should return a `Netinfo` type")
	assert.IsType(t, expected.up, actual.up, "`getNetinfo()` should return a `up` field with a float64 type")
	assert.IsType(t, expected.down, actual.down, "`getNetinfo()` should return a `down` field with a float64 type")
}

// TestNetIocounters test if `netIocounters()` return a value with a []net.IOCountersStat slice
func TestNetIocounters(t *testing.T) {
	expected := []net.IOCountersStat{}
	actual, _ := netIocounters(false)

	assert.IsType(t, expected, actual, "`netIocounters()` should return a []net.IOCountersStat slice")
}

//TODO add tests for errors --> must return empty/zero values
