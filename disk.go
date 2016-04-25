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
	var ret []diskinfo

	// get the partitions list
	partitions, err := diskPartitions(false) //`false` to get only physical disks
	if err == nil {
		// get usage stats for each partitions
		for i := range partitions {
			disk, err := diskUsage(partitions[i].Mountpoint) //TODO rename `err` variables names to avoid confusion ?
			if err == nil {
				d := diskinfo{
					device:      partitions[i].Device,
					path:        disk.Path,
					total:       strconv.FormatFloat(float64(disk.Total)/(1024*1024*1024), 'f', 2, 64), // (1024*1024*1024) to convert to GiB from `gopesutil`
					used:        strconv.FormatFloat(float64(disk.Used)/(1024*1024*1024), 'f', 2, 64),
					usedPercent: int(disk.UsedPercent),
				}
				ret = append(ret, d)
			}
		}
	}
	return ret
}

// wrap `disk.Partitions()` in an unexported variable for testability
var diskPartitions = func(all bool) ([]disk.PartitionStat, error) {
	return disk.Partitions(all)
}

// wrap `disk.Usage()` in an unexported variable for testability
var diskUsage = func(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}
