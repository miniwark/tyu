package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetBIOSStat test the returned fields values and types of `getBIOSStat()`
func TestGetBIOSStat(t *testing.T) {
	// setup the faking of readAndTrimFile
	oldReadAndTrimFile := readAndTrimFile
	readAndTrimFile = func(path string) (ret string, err error) {
		return "abc", nil
	}

	// test
	expected := biosStat{
		boardName:   "abc",
		boardVendor: "abc",
		biosVendor:  "abc",
		biosVersion: "abc",
		biosDate:    "abc",
	}
	actual, err := getBIOSStat()

	assert.NoError(t, err, "`getBIOSStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getBIOSStat()` should be equal to main.biosStat{boardName:\"abc\", boardVendor:\"abc\", biosVendor:\"abc\", biosVersion:\"abc\", biosBate:\"abc\"}")

	// teardown
	readAndTrimFile = oldReadAndTrimFile
}

// TestGetDiskStatErrorCase1 test than getCPUStat() transmit the error from readAndTrimFile()
func TestGetBIOSStatErrorCase1(t *testing.T) {
	// setup the faking of readAndTrimFile
	oldReadAndTrimFile := readAndTrimFile
	readAndTrimFile = func(path string) (ret string, err error) {
		ret = ""
		err = errors.New("Error 1")
		return ret, err
	}

	//test
	expected := errors.New("Error 1; Error 1; Error 1; Error 1; Error 1") //error repeted 5 time because readAndTrimFile() invoked 5 times in getBIOSStat()
	_, actual := getBIOSStat()

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`getBIOSStat()` should be an error equal to \"Error 1; Error 1; Error 1; Error 1; Error 1\"")

	// teardown
	readAndTrimFile = oldReadAndTrimFile
}

// TestGetBIOSStatType test if `getBIOSStat()` return a `biosStat` type and if each fields have the correct types
// Types regression testing
func TestGetBIOSStatType(t *testing.T) {
	expected := biosStat{
		boardName:   "", // the result values of the fields are not tested
		boardVendor: "",
		biosVendor:  "",
		biosVersion: "",
		biosDate:    "",
	}
	actual, _ := getBIOSStat()
	assert.IsType(t, expected, actual, "`getBIOSStat()` should return a `biosStat` type")
	assert.IsType(t, expected.boardName, actual.boardName, "`getBIOSStat()` should return a `boardName` field with a string type")
	assert.IsType(t, expected.boardVendor, actual.boardVendor, "`getBIOSStat()` should return a `boardVendor` field with a string type")
	assert.IsType(t, expected.biosVendor, actual.biosVendor, "`getBIOSStat()` should return a `biosVendor` field with a string type")
	assert.IsType(t, expected.biosVersion, actual.biosVersion, "`getBIOSStat()` should return a `biosVersion` field with a string type")
	assert.IsType(t, expected.biosDate, actual.biosDate, "`getBIOSStat()` should return a `biosDate` field with a string type")
}
