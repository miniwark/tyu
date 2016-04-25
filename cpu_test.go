package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

// TestGetCPUinfo test the returned fields values of `getCPUinfo()`
func TestGetCPUinfo(t *testing.T) {
	// setup the faking of `cpu.Info()`
	oldcpuInfo := cpuInfo //TODO change variable name to avoid confuion with `cpuinfo` struct (or rename the struct)
	cpuInfo = func() ([]cpu.InfoStat, error) {
		ret := []cpu.InfoStat{
			{
				VendorID:  string("vendor"),
				ModelName: string("model"),
				Mhz:       float64(100),
			},
			{
				VendorID:  string("vendor"), // two CPUs --> cpuinfo.count = "2"
				ModelName: string("model"),
				Mhz:       float64(100),
			},
		}
		return ret, nil
	}

	// test
	expected := cpuinfo{
		count:     strconv.FormatInt(2, 10),
		vendorID:  "vendor",
		modelName: "model",
		cpuMhz:    strconv.FormatInt(100, 10),
	}
	actual := getCPUinfo()

	assert.Equal(t, expected, actual, "`getCPUinfo()` should be equal to --> cpuinfo{count:\"2\", vendorID:\"vendor\", modelName:\"model\", cpuMhz:\"100\"}")
	// teardown
	cpuInfo = oldcpuInfo
}

// TestGetCPUinfoType test if `getCPUinfo()` return a `cpuinfo` type and if each fields
// of `cpuinfo` have the correct types
// Types regression testing
func TestGetCPUinfoType(t *testing.T) {
	expected := cpuinfo{
		count:     "", // the result values of the fields are not tested
		vendorID:  "",
		modelName: "",
		cpuMhz:    "",
	}
	actual := getCPUinfo()

	assert.IsType(t, expected, actual, "`getCPUinfo()` should return a `cpuinfo` type")
	assert.IsType(t, expected.count, actual.count, "`getCPUinfo()` should return a `count` field with a string type")
	assert.IsType(t, expected.vendorID, actual.vendorID, "`getCPUinfo()` should return a `vendorID` field with a string type")
	assert.IsType(t, expected.modelName, actual.modelName, "`getCPUinfo()` should return a `modelName` field with a string type")
	assert.IsType(t, expected.cpuMhz, actual.cpuMhz, "`getCPUinfo()` should return a `cpuMhz` field with a string type")
}

// TestGetCPUpercent test the returned value of `getCPUpercent()`
func TestGetCPUpercent(t *testing.T) {
	// setup the faking of `cpu.Percent()`
	oldcpuPercent := cpuPercent
	cpuPercent = func(interval time.Duration, percpu bool) ([]float64, error) {
		ret := []float64{100}
		return ret, nil
	}
	// test
	expected := 100
	actual := getCPUpercent()
	assert.Equal(t, expected, actual, "`getCPUpercent` should be equal to --> 100")

	// teardown
	cpuPercent = oldcpuPercent
}

// TestGetCPUinfoType test if `getCPUinfo()` return a value with a int type
func TestGetCPUpercentType(t *testing.T) {
	expected := int(0) // the result value is not tested
	actual := getCPUpercent()
	assert.IsType(t, expected, actual, "`getCPUpercent()` should return an `int`")
}
