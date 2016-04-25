package main

import (
	"strconv"

	"github.com/shirou/gopsutil/process"
)

// procinfo represent the processes informations
type procinfo struct {
	total   string
	running string
	//TODO add zombies threads ?
}

// Get informations about the processes
func getProcinfo() procinfo {
	ret := procinfo{}

	pids, err := processPids() //TODO replace by something like psutil.process_iter() if available in gopsutils
	if err == nil {
		run := 0
		for i := range pids {
			proc, err := processNewProcess(pids[i])
			if err == nil { //TODO rename `err` variables names to avoid confusion ?
				status, err := proc.Status()
				if err == nil {
					if status == "R" { // "R" for running process
						run++
					}
				}
			}
		}
		ret.total = strconv.Itoa(len(pids))
		ret.running = strconv.Itoa(run)
	}
	return ret
}

// wrap `process.Pids()` in an unexported variable for testability
var processPids = func() ([]int32, error) {
	return process.Pids()
}

// wrap `process.NewProcess()` in an unexported variable for testability
var processNewProcess = func(pid int32) (*process.Process, error) {
	return process.NewProcess(pid)
}
