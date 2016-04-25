package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInt8SliceToString test the returned type and value of `int8SliceToString()`
func TestInt8SliceToString(t *testing.T) {
	expected := "123"
	actual := int8SliceToString([]int8{49, 50, 51})

	assert.IsType(t, expected, actual, "`getNetinfo()` should return a `string`")
	assert.Equal(t, expected, actual, "`int8SliceToString([]int8{49, 50, 51})` should be equal to --> \"123\"")
}

// TestReadAndTrimFile test the returned type and value of `readAndTrimFile()`
// func TestReadAndTrimFile(t *testing.T) {
// }

// TODO maybe stick to the testing stdlib intead of testify
// func TestInt8SliceToString(t *testing.T) {
// 	char := []int8{49, 50, 51} // See the ASCII table: 49 = "1" ; 50 ="2" ; 51 = "3"
// 	expected := "123"
// 	actual := int8SliceToString(char)
// 	if actual != expected {
// 		t.Error("expected:", expected)
// 		t.Error("actual  :", actual)
// 	}
// }
