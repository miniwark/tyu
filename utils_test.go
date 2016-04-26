package main

import (
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
	actual := readAndTrimFile("utils_test.go")

	assert.IsType(t, expected, actual, "`readAndTrimFile()` should return a `string`")
	assert.Regexp(t, regexp.MustCompile(expected), actual, "`readAndTrimFile(\"utils_test.go\")` should contain \"package main\"")
	//TODO improve tests with Equal instead of Regexp by using a temp file ?
}

// TODO maybe stick to the testing stdlib intead of testify
// func TestInt8SliceToString(t *testing.T) {
// 	char := []int8{49, 50, 51}
// 	expected := "123"
// 	actual := int8SliceToString(char)
// 	if actual != expected {
// 		t.Error("expected:", expected)
// 		t.Error("actual  :", actual)
// 	}
// }
