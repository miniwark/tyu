package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBIOSinfo(t *testing.T) {
	expected := biosinfo{
		boardName:   "abc",
		boardVendor: "abc",
		biosVendor:  "abc",
		biosVersion: "abc",
		biosDate:    "abc",
	}
	actual := getBIOSinfo()
	assert.IsType(t, expected, actual, "`getBIOSinfo` should return a `biosinfo` type")
}
