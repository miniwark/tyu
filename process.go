package main

import (
	"strconv"

	"github.com/shirou/gopsutil/process"
)

// procinfo represent the processes informations
type procinfo struct {
	total   string
	running string
	//TODO sleeping and/or threads ?
}

// Get informations about the processes
func getProcinfo() procinfo {
	pids, err := process.Pids() //TODO replace by somethin like psutil.process_iter()
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	var run int

	for i := range pids {
		proc, err := process.NewProcess(pids[i])
		if err != nil {
			panic(err) //TODO do not panic but manage the error
		}
		status, err := proc.Status()
		if err != nil {
			panic(err) //TODO do not panic but manage the error
		}
		if status == "R" {
			run++
		}
	}
	// the status value is not the same than psutils, the status comme directly from /proc/[pid]/stat
	// with one character from the string "RSDZTW" where R is running, S is sleeping in an
	// interruptible wait, D is waiting in uninterruptible disk sleep, Z is zombie,
	// T is traced or stopped (on a signal), and W is paging.

	ret := procinfo{
		total:   strconv.Itoa(len(pids)),
		running: strconv.Itoa(run),
	}
	return ret
}
