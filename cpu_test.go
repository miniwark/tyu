package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCPUinfo(t *testing.T) {
	expected := cpuinfo{
		count:     string(123),
		vendorID:  "abc",
		modelName: "abc",
		cpuMhz:    string(123),
	}
	actual := getCPUinfo()
	assert.IsType(t, expected, actual, "`getCPUinfo()` should return a `cpuinfo` type")
}

func TestGetCPUpercent(t *testing.T) {
	expected := int(100)
	actual := getCPUpercent()
	assert.IsType(t, expected, actual, "`getCPUpercent()` should return an `int`")
}
