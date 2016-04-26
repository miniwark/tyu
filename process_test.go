package main

import (
	"strconv"
	"testing"

	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

// TestGetProcinfo test the returned fields values and types of `getProcinfo()`
func TestGetProcinfo(t *testing.T) {
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
	expected := procinfo{
		total:   strconv.FormatInt(1, 10),
		running: strconv.FormatInt(1, 10),
	}
	actual := getProcinfo()
	assert.Equal(t, expected, actual, "`getProcinfo` should be equal to main.procinfo{total:\"1\", running:\"1\"}")

	// teardown
	processPids = oldProcessPids
	procStatus = oldProcStatus
}

// TestGetProcinfoType test if `getProcinfo()` return a `procinfo` type and if each fields have the correct types
func TestGetProcinfoType(t *testing.T) {
	expected := procinfo{
		total:   "", // the result values of the fields are not tested
		running: "",
	}
	actual := getProcinfo()

	assert.IsType(t, expected, actual, "`getProcinfo` should return a `procinfo` type")
	assert.IsType(t, expected.total, actual.total, "`getMeminfo()` should return a `total` field with a string type")
	assert.IsType(t, expected.running, actual.running, "`getMeminfo()` should return a `ramUsed` field with a string type")
}
