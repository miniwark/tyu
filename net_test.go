package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNetinfo(t *testing.T) {
	expected := Netinfo{
		up:   float64(123.123),
		down: float64(123.123),
	}
	actual := getNetinfo()
	assert.IsType(t, expected, actual, "`getNetinfo()` should return a `Netinfo` type")
	//assert.Equal(t, expected, actual, "`getNetinfo()` should be equal to --> Netinfo{up: 123.123, down: 123.123}")
	//TODO add mock or monckey patching to test equality
}
