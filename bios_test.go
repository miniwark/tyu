package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetBIOSinfoType test if `getBIOSinfo()` return a `biosinfo` type and if each fields have the correct types
// Types regression testing
func TestGetBIOSinfoType(t *testing.T) {
	expected := biosinfo{
		boardName:   "", // the result values of the fields are not tested
		boardVendor: "",
		biosVendor:  "",
		biosVersion: "",
		biosDate:    "",
	}
	actual, _ := getBIOSinfo()
	assert.IsType(t, expected, actual, "`getBIOSinfo()` should return a `biosinfo` type")
	assert.IsType(t, expected.boardName, actual.boardName, "`getBIOSinfo()` should return a `boardName` field with a string type")
	assert.IsType(t, expected.boardVendor, actual.boardVendor, "`getBIOSinfo()` should return a `boardVendor` field with a string type")
	assert.IsType(t, expected.biosVendor, actual.biosVendor, "`getBIOSinfo()` should return a `biosVendor` field with a string type")
	assert.IsType(t, expected.biosVersion, actual.biosVersion, "`getBIOSinfo()` should return a `biosVersion` field with a string type")
	assert.IsType(t, expected.biosDate, actual.biosDate, "`getBIOSinfo()` should return a `biosDate` field with a string type")
}
