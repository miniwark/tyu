package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDiskinfo(t *testing.T) {
	expected := []diskinfo{
		{
			device:      "abc",
			path:        "abc",
			total:       string(123),
			used:        string(123),
			usedPercent: int(100),
		},
		{
			device:      "abc",
			path:        "abc",
			total:       string(123),
			used:        string(123),
			usedPercent: int(100),
		},
	}
	actual := getDiskinfo()
	assert.IsType(t, expected, actual, "`getDiskinfo()` should return a `[]diskinfo` slice")
}
