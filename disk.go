package main

import (
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

// diskinfo represents the physical disk usage statistics
type diskinfo struct {
	device      string // ex. '/dev/sda1'
	path        string // ex. '/''
	total       string
	used        string
	usedPercent int
}

// Return a slice of physical disks usage statistics by using `gopesutil` package
func getDiskinfo() []diskinfo {
	// get the partitions list
	partitions, err := disk.Partitions(false) //`false` to get only physical disks
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	var ret []diskinfo

	// get usage stats for each partitions
	for i := range partitions {
		disk, err := disk.Usage(partitions[i].Mountpoint)
		if err != nil {
			panic(err) //TODO do not panic but manage the error
		}
		d := diskinfo{
			device:      partitions[i].Device,
			path:        disk.Path,
			total:       strconv.FormatFloat(float64(disk.Total)/(1024*1024*1024), 'f', 2, 64), // (1024*1024*1024) to convert to GiB from `gopesutil`
			used:        strconv.FormatFloat(float64(disk.Used)/(1024*1024*1024), 'f', 2, 64),
			usedPercent: int(disk.UsedPercent),
		}
		ret = append(ret, d)
	}
	return ret
}
