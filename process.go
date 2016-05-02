package main

import (
	"strconv"

	"github.com/shirou/gopsutil/process"
)

// procStat represent the processes informations
type procStat struct {
	total   string
	running string
	//TODO add zombies threads ?
}

// getProcStat get informations about the processes
func getProcStat() (ret procStat, err error) {
	pids, err := processPids() //TODO replace by something like psutil.process_iter() if available in gopsutils
	if err == nil {
		run := 0
		for i := range pids {
			proc, err1 := processNewProcess(pids[i])
			if err1 == nil { //TODO
				status, err2 := procStatus(proc)
				if err2 == nil {
					if status == "R" { // "R" for running process
						run++
					}
				} else {
					err = appendError(err, err2)
				}
			} else {
				err = appendError(err, err1)
			}
		}
		ret.total = strconv.Itoa(len(pids))
		ret.running = strconv.Itoa(run)
	}
	return ret, err
}

// wrap `process.Pids()` in an unexported variable for testability
var processPids = func() ([]int32, error) {
	return process.Pids()
}

// wrap `process.NewProcess()` in an unexported variable for testability
var processNewProcess = func(pid int32) (*process.Process, error) {
	return process.NewProcess(pid)
}

// wrap `Process.Status()` in an unexported variable for testability
var procStatus = func(proc *process.Process) (string, error) {
	return proc.Status()
}
