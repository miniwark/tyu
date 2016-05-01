package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetBIOSinfoType test if `getBIOSStat()` return a `biosStat` type and if each fields have the correct types
// Types regression testing
func TestGetBIOSinfoType(t *testing.T) {
	expected := biosStat{
		boardName:   "", // the result values of the fields are not tested
		boardVendor: "",
		biosVendor:  "",
		biosVersion: "",
		biosDate:    "",
	}
	actual, _ := getBIOSStat()
	assert.IsType(t, expected, actual, "`getBIOSStat()` should return a `biosinfo` type")
	assert.IsType(t, expected.boardName, actual.boardName, "`getBIOSStat()` should return a `boardName` field with a string type")
	assert.IsType(t, expected.boardVendor, actual.boardVendor, "`getBIOSStat()` should return a `boardVendor` field with a string type")
	assert.IsType(t, expected.biosVendor, actual.biosVendor, "`getBIOSStat()` should return a `biosVendor` field with a string type")
	assert.IsType(t, expected.biosVersion, actual.biosVersion, "`getBIOSStat()` should return a `biosVersion` field with a string type")
	assert.IsType(t, expected.biosDate, actual.biosDate, "`getBIOSStat()` should return a `biosDate` field with a string type")
}
