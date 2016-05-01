package main

import (
	"testing"

	"github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

// TestGetDiskStat test the returned fields values and types of `getDiskStat()`
func TestGetDiskStat(t *testing.T) {
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
	expected := []diskStat{
		{
			device:      "/dev/device",
			path:        "/path",
			total:       "1.00",
			used:        "1.00",
			usedPercent: int(100),
		},
	}
	actual, err := getDiskStat()

	assert.NoError(t, err, "`getDiskStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getDiskStat()` should be equal to []main.diskStat{main.diskStat{device:\"/dev/device\", path:\"/path\", total:\"1.00\", used:\"1.00\", usedPercent:100}}")

	// teardown
	diskPartitions = oldDiskPartitions
	diskUsage = oldDiskUsage
}

// TestGetDiskStatType test if `getDiskStat()` return a value with a []diskStat` slice
func TestGetDiskStatType(t *testing.T) {
	expected := []diskStat{
		{
			device:      "", // the result value is not tested
			path:        "",
			total:       "",
			used:        "",
			usedPercent: int(0),
		},
	}
	actual, _ := getDiskStat()

	assert.IsType(t, expected[0].device, actual[0].device, "`getDiskStat()` should return a `device` field with a string type")
	assert.IsType(t, expected[0].path, actual[0].path, "`getDiskStat()` should return a `path` field with a string type")
	assert.IsType(t, expected[0].total, actual[0].total, "`getDiskStat()` should return a `total` field with a string type")
	assert.IsType(t, expected[0].used, actual[0].used, "`getDiskStat()` should return a `used` field with a string type")
	assert.IsType(t, expected[0].usedPercent, actual[0].usedPercent, "`getDiskStat()` should return a `usedPercent` field with an int type")
	assert.IsType(t, expected, actual, "`getDiskStat()` should return a `[]main.diskStat` slice")
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
