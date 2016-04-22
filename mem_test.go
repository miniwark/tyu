package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeminfo(t *testing.T) {
	expected := meminfo{
		ramTotal:        string(123),
		ramUsed:         string(123),
		ramUsedPercent:  int(100),
		swapTotal:       string(123),
		swapUsed:        string(123),
		swapUsedPercent: int(100),
	}
	actual := getMeminfo()
	assert.IsType(t, expected, actual, "`getMeminfo()` should return a `meminfo` type")
}
