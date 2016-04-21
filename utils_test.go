package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestInt8SliceToString(t *testing.T) {
	char := []int8{49, 50, 51}
	assert.Equal(t, "123", int8SliceToString(char), "`int8SliceToString([]int8{49, 50, 51})` should be equal to \"123\"")
}
