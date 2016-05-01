package main

import (
	"testing"

	"github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

// TestGetDiskinfo test the returned fields values and types of `getDiskinfo()`
func TestGetDiskinfo(t *testing.T) {
	// setup the faking of `disk.Partitions()` & `disk.Usage()`
	oldDiskPartitions := diskPartitions
	oldDiskUsage := diskUsage
	diskPartitions = func(all bool) ([]disk.PartitionStat, error) {
		ret := []disk.PartitionStat{
			{
				Device:     "/dev/device",
				Mountpoint: "/mount/point",
			},
		}
		return ret, nil
	}
	diskUsage = func(path string) (*disk.UsageStat, error) {
		ret := &disk.UsageStat{
			Path:        "/path",
			Total:       uint64(1024 * 1024 * 1024), // KiB to GiB conversion is implicitly tested --> ((1024 * 1024 * 1024) / (1024 * 1024 * 1024)) = 1.00
			Used:        uint64(1024 * 1024 * 1024),
			UsedPercent: float64(100),
		}
		return ret, nil
	}

	// test
	expected := []diskinfo{
		{
			device:      "/dev/device",
			path:        "/path",
			total:       "1.00",
			used:        "1.00",
			usedPercent: int(100),
		},
	}
	actual := getDiskinfo()

	assert.Equal(t, expected, actual, "`getDiskinfo` should be equal to []main.diskinfo{main.diskinfo{device:\"/dev/device\", path:\"/path\", total:\"1.00\", used:\"1.00\", usedPercent:100}}")

	// teardown
	diskPartitions = oldDiskPartitions
	diskUsage = oldDiskUsage
}

// TestGetDiskinfoType test if `getDiskinfo()` return a value with a []diskinfo` slice
func TestGetDiskinfoType(t *testing.T) {
	expected := []diskinfo{
		{
			device:      "", // the result value is not tested
			path:        "",
			total:       "",
			used:        "",
			usedPercent: int(0),
		},
	}
	actual := getDiskinfo()

	assert.IsType(t, expected[0].device, actual[0].device, "`getDiskinfo()` should return a `device` field with a string type")
	assert.IsType(t, expected[0].path, actual[0].path, "`getDiskinfo()` should return a `path` field with a string type")
	assert.IsType(t, expected[0].total, actual[0].total, "`getDiskinfo()` should return a `total` field with a string type")
	assert.IsType(t, expected[0].used, actual[0].used, "`getDiskinfo()` should return a `used` field with a string type")
	assert.IsType(t, expected[0].usedPercent, actual[0].usedPercent, "`getDiskinfo()` should return a `usedPercent` field with an int type")
	assert.IsType(t, expected, actual, "`getDiskinfo()` should return a `[]main.diskinfo` slice")
}

// TestDiskPartitions test if `diskPartitions()` return a value with a []disk.PartitionStat slice
func TestDiskPartitions(t *testing.T) {
	expected := []disk.PartitionStat{}
	actual, _ := diskPartitions(false)

	assert.IsType(t, expected, actual, "`diskPartitions()` should return a []disk.PartitionStat slice")
}

// TestDiskUsage test if `diskUsage()` return a value with a *disk.UsageStat type
func TestDiskUsage(t *testing.T) {
	expected := &disk.UsageStat{}
	actual, _ := diskUsage("")

	assert.IsType(t, expected, actual, "`diskUsage()` should return a *disk.UsageStat type")
}
