package main

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInt8SliceToString test the returned type and value of `int8SliceToString()`
func TestInt8SliceToString(t *testing.T) {
	expected := "abc"
	actual := int8SliceToString([]int8{97, 98, 99}) // See the ASCII table: 97 = "a" ; 98 ="b" ; 99 = "c"

	assert.IsType(t, expected, actual, "`int8SliceToString()` should return a `string`")
	assert.Equal(t, expected, actual, "`int8SliceToString([]int8{97, 98, 99})` should be equal to \"abc\"")
}

// TestReadAndTrimFile test the returned type and value of `readAndTrimFile()`
func TestReadAndTrimFile(t *testing.T) {
	expected := "package main"
	actual, err := readAndTrimFile("utils_test.go")

	assert.NoError(t, err, "`readAndTrimFile()` should not have returned an error")
	assert.IsType(t, expected, actual, "`readAndTrimFile()` should return a `string`")
	assert.Regexp(t, regexp.MustCompile(expected), actual, "`readAndTrimFile(\"utils_test.go\")` should contain \"package main\"")
	//TODO improve tests with Equal instead of Regexp by using a temp file ?
}

// TestReadAndTrimFileErrorCase1 test than readAndTrimFile() transmit the error from ioutil.ReadFile()
func TestReadAndTrimFileErrorCase1(t *testing.T) {
	// setup the faking of `ioutil.ReadFile()`
	oldIoutilReadFile := ioutilReadFile
	ioutilReadFile = func(filename string) ([]byte, error) {
		ret := []byte{}
		err := errors.New("Error 1")
		return ret, err
	}

	// test
	expected := errors.New("Error 1")
	_, actual := readAndTrimFile("/fakepath")

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`readAndTrimFile()` should be an error equal to \"Error 1\"")

	// teardown
	ioutilReadFile = oldIoutilReadFile
}

// TestAppendErrorCase1 tests the case where appendError() was provided with two errors as arguments
func TestAppendErrorCase1(t *testing.T) {
	err1 := errors.New("Error 1")
	err2 := errors.New("Error 2")

	expected := errors.New("Error 1; Error 2")
	actual := appendError(err1, err2)

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`appendError(errors.New(\"Error 1\"), errors.New(\"Error 2\"))` should be an error equal to \"Error 1; Error 2\"")
}

// TestAppendErrorCase2 tests the case where appendError() was provided with one error in the first arguments
func TestAppendErrorCase2(t *testing.T) {
	err1 := errors.New("Error 1")

	expected := errors.New("Error 1")
	actual := appendError(err1, nil)

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`appendError(errors.New(\"Error 1\"), nil)` should be an error equal to \"Error 1\"")
}

// TestAppendErrorCase3 tests the case where appendError() was provided with one error in the seconds arguments
func TestAppendErrorCase3(t *testing.T) {
	err2 := errors.New("Error 2")

	expected := errors.New("Error 2")
	actual := appendError(nil, err2)

	assert.EqualError(t, expected, fmt.Sprintf("%v", actual), "`appendError(nil, errors.New(\"Error 2\"))` should be an error equal to \"Error 2\"")
}

// TestAppendErrorNil tests the case where appendError() was provided with nil errors as arguments
func TestAppendErrorNil(t *testing.T) {
	err := appendError(nil, nil)

	assert.NoError(t, err, "`appendError(nil, nil)` should not be an error")
}

// TestIoutilReadFile test if `ioutilReadFile()` return a value with a []byte type
func TestIoutilReadFile(t *testing.T) {
	expected := []byte{} // the result value is not tested
	actual, _ := ioutilReadFile("/fakepath")

	assert.IsType(t, expected, actual, "`ioutilReadFile()` should return a []byte type")
}
