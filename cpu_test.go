package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

// TestGetCPUStat test the returned fields values of `getCPUStat()`
func TestGetCPUStat(t *testing.T) {
	// setup the faking of `cpu.Info()`
	oldcpuInfo := cpuInfo
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
	expected := cpuStat{
		count:     strconv.FormatInt(2, 10),
		vendorID:  "vendor",
		modelName: "model",
		mhz:       strconv.FormatInt(100, 10),
	}
	actual, err := getCPUStat()

	assert.NoError(t, err, "`getCPUStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getCPUStat()` should be equal to main.cpuStat{count:\"2\", vendorID:\"vendor\", modelName:\"model\", mhz:\"100\"}")

	// teardown
	cpuInfo = oldcpuInfo
}

// TestGetCPUStatErrorCase1 test than getCPUStat() transmit the error from cpu.Info()
func TestGetCPUStatErrorCase1(t *testing.T) {
	// setup the faking of `cpu.Info()`
	oldcpuInfo := cpuInfo
	cpuInfo = func() ([]cpu.InfoStat, error) {
		err := errors.New("Error 1")
		return nil, err
	}

	//test
	expected := errors.New("Error 1")
	_, actual := getCPUStat()
	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getCPUStat()` should be an error equal to \"Error 1\"")

	// teardown
	cpuInfo = oldcpuInfo
}

// TestGetCPUStatType test if `getCPUStat()` return a `cpuStat` type and if each fields have the correct types
func TestGetCPUStatType(t *testing.T) {
	expected := cpuStat{
		count:     "", // the result values of the fields are not tested
		vendorID:  "",
		modelName: "",
		mhz:       "",
	}
	actual, _ := getCPUStat()

	assert.IsType(t, expected, actual, "`getCPUStat()` should return a `cpuStat` type")
	assert.IsType(t, expected.count, actual.count, "`getCPUStat()` should return a `count` field with a string type")
	assert.IsType(t, expected.vendorID, actual.vendorID, "`getCPUStat()` should return a `vendorID` field with a string type")
	assert.IsType(t, expected.modelName, actual.modelName, "`getCPUStat()` should return a `modelName` field with a string type")
	assert.IsType(t, expected.mhz, actual.mhz, "`getCPUStat()` should return a `mhz` field with a string type")
}

// TestGetCPUPercent test the returned value of `getCPUPercent()`
func TestGetCPUPercent(t *testing.T) {
	// setup the faking of `cpu.Percent()`
	oldcpuPercent := cpuPercent
	cpuPercent = func(interval time.Duration, percpu bool) ([]float64, error) {
		ret := []float64{100}
		return ret, nil
	}

	// test
	expected := 100
	actual, err := getCPUPercent()

	assert.NoError(t, err, "`getCPUPercent()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getCPUPercent` should be equal to --> 100")

	// teardown
	cpuPercent = oldcpuPercent
}

// TestGetCPUPercentErrorCase1 test than getCPUPercent() transmit the error from cpu.Percent()
func TestGetCPUPercentErrorCase1(t *testing.T) {
	// setup the faking of `cpu.Percent()`
	oldcpuPercent := cpuPercent
	cpuPercent = func(interval time.Duration, percpu bool) ([]float64, error) {
		err := errors.New("Error 1")
		return nil, err
	}

	//test
	expected := errors.New("Error 1")
	_, actual := getCPUPercent()
	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getCPUPercent()` should be an error equal to \"Error 1\"")

	// teardown
	cpuPercent = oldcpuPercent
}

// TestGetCPUPercentType test if `getCPUPercent()` return a value with a int type
func TestGetCPUPercentType(t *testing.T) {
	expected := int(0) // the result value is not tested
	actual, _ := getCPUPercent()

	assert.IsType(t, expected, actual, "`getCPUPercent()` should return an `int`")
}

// TestCpuInfo test if `cpu.Info()` return a value with a []cpu.InfoStat slice
func TestCpuInfo(t *testing.T) {
	expected := []cpu.InfoStat{}
	actual, _ := cpu.Info()

	assert.IsType(t, expected, actual, "`cpuInfo()` should return a []cpu.InfoStat slice")
}

// TestCpuPercent test if `cpu.Percent()` return a value with a []float64 slice
func TestCpuPercent(t *testing.T) {
	expected := []float64{}
	actual, _ := cpuPercent((500 * time.Millisecond), false)

	assert.IsType(t, expected, actual, "`cpuPercent()` should return a []float64 slice")
}
