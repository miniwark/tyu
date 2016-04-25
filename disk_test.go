package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetDiskinfoType test if `getDiskinfo()` return `[]diskinfo` slice" and if each fields have the correct types
// Types regression testing
func TestGetDiskinfoType(t *testing.T) {
	expected := []diskinfo{
		{
			device:      "", // the result values of the fields are not tested
			path:        "",
			total:       "",
			used:        "",
			usedPercent: int(0),
		},
	}
	actual := getDiskinfo()
	assert.IsType(t, expected, actual, "`getDiskinfo()` should return a `[]diskinfo` slice")
	assert.IsType(t, expected[0].device, actual[0].device, "`getDiskinfo()` should return a `device` field with a string type")
	assert.IsType(t, expected[0].path, actual[0].path, "`getDiskinfo()` should return a `path` field with a string type")
	assert.IsType(t, expected[0].total, actual[0].total, "`getDiskinfo()` should return a `total` field with a string type")
	assert.IsType(t, expected[0].used, actual[0].used, "`getDiskinfo()` should return a `used` field with a string type")
	assert.IsType(t, expected[0].usedPercent, actual[0].usedPercent, "`getDiskinfo()` should return a `usedPercent` field with an int type")
}
