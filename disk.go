package main

import (
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

// `diskinfo` represents the physical disk usage statistics
type diskinfo struct {
	device      string // ex. '/dev/sda1'
	path        string // ex. '/''
	total       string
	used        string
	free        string
	usedPercent int
}

// Return a slice of physical disks usage statistics by using `gopesutil` package
func getDiskinfo() []diskinfo {
	// get the partitions list
	partitions, err := disk.Partitions(false) //TODO understand the true/false
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	var result []diskinfo

	// get usage stats for each partitions
	for i := range partitions {
		disk, err := disk.Usage(partitions[i].Mountpoint)
		if err != nil {
			panic(err) //TODO do not panic but manage the error
		}
		if strings.HasPrefix(partitions[i].Device, "/dev") { // only look for physical disks //TODO must improve work only on Linux for now
			d := diskinfo{
				device:      partitions[i].Device,
				path:        disk.Path,
				total:       strconv.FormatUint(disk.Total, 10),
				used:        strconv.FormatUint(disk.Free, 10),
				free:        strconv.FormatUint(disk.Used, 10),
				usedPercent: int(disk.UsedPercent),
			}
			result = append(result, d)
		}
	}
	return result
}
