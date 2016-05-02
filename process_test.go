package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

// TestGetProcStat test the returned fields values and types of `getProcStat()`
func TestGetProcStat(t *testing.T) {
	// setup the faking of `process.Pids()` & `process.NewProcess()`
	oldProcessPids := processPids
	oldProcStatus := procStatus
	processPids = func() ([]int32, error) {
		ret := []int32{1} // one fake Pid with number 1
		return ret, nil
	}
	procStatus = func(proc *process.Process) (string, error) {
		return "R", nil // all processes have the "R" status
	}

	// test
	expected := procStat{
		total:   strconv.FormatInt(1, 10),
		running: strconv.FormatInt(1, 10),
	}
	actual, err := getProcStat()

	assert.NoError(t, err, "`getProcStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getProcStat()` should be equal to main.procStat{total:\"1\", running:\"1\"}")

	// teardown
	processPids = oldProcessPids
	procStatus = oldProcStatus
}

// TestGetDiskStatErrorCase1 test than getProcStat() transmit the error from process.Pids()
func TestGetProcStatErrorCase1(t *testing.T) {
	// setup the faking of `process.Pids()`
	oldProcessPids := processPids
	processPids = func() ([]int32, error) {
		err := errors.New("Error 1")
		return nil, err
	}

	// test
	expected := errors.New("Error 1")
	_, actual := getProcStat()

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getProcStat()` should be an error equal to \"Error 1\"")

	// teardown
	processPids = oldProcessPids
}

// TestGetDiskStatErrorCase2 test than getProcStat() transmit the error from process.NewProcess()
func TestGetProcStatErrorCase2(t *testing.T) {
	// setup the faking of `process.Pids()` & process.NewProcess()
	oldProcessPids := processPids
	oldProcessNewProcess := processNewProcess
	processPids = func() ([]int32, error) {
		ret := []int32{1} // one fake Pid with number 1
		return ret, nil
	}
	processNewProcess = func(pid int32) (*process.Process, error) {
		err := errors.New("Error 2")
		return nil, err
	}

	// test
	expected := errors.New("Error 2")
	_, actual := getProcStat()

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getProcStat()` should be an error equal to \"Error 2\"")

	// teardown
	processPids = oldProcessPids
	processNewProcess = oldProcessNewProcess
}

// TestGetProcStatErrorCase3 test than getProcStat() transmit the error from procStatus()
func TestGetProcStatErrorCase3(t *testing.T) {
	// setup the faking of `process.Pids()` & `process.NewProcess()`
	oldProcessPids := processPids
	oldProcStatus := procStatus
	processPids = func() ([]int32, error) {
		ret := []int32{1} // one fake Pid with number 1
		return ret, nil
	}
	procStatus = func(proc *process.Process) (string, error) {
		ret := ""
		err := errors.New("Error 3")
		return ret, err
	}

	// test
	expected := errors.New("Error 3")
	_, actual := getProcStat()

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getProcStat()` should be an error equal to \"Error 3\"")

	// teardown
	processPids = oldProcessPids
	procStatus = oldProcStatus
}

// TestGetProcStatType test if `getProcStat()` return a `procStat` type and if each fields have the correct types
func TestGetProcStatType(t *testing.T) {
	expected := procStat{
		total:   "", // the result values of the fields are not tested
		running: "",
	}
	actual, _ := getProcStat()

	assert.IsType(t, expected, actual, "`getProcStat()` should return a `procStat` type")
	assert.IsType(t, expected.total, actual.total, "`getProcStat()` should return a `total` field with a string type")
	assert.IsType(t, expected.running, actual.running, "`getProcStat()` should return a `ramUsed` field with a string type")
}

// TestProcessPids test if `processPids()` return a []int32 slice
func TestProcessPids(t *testing.T) {
	expected := []int32{}
	actual, _ := processPids()

	assert.IsType(t, expected, actual, "`processPids()` should return a []int32 slice")
}

// TestProcessNewProcess test if `processPids()` return a *process.Process type
func TestProcessNewProcess(t *testing.T) {
	expected := &process.Process{}
	actual, _ := processNewProcess(1) //pid nimber is not tested

	assert.IsType(t, expected, actual, "`processNewProcess()` should return a *process.Process type")
}

// TestProcStatus test if `procStatus()` return a value with a string type
func TestProcStatus(t *testing.T) {
	expected := "" // the result value is not tested
	actual, _ := procStatus(&process.Process{})

	assert.IsType(t, expected, actual, "`procStatus()` should return a string type")
}
