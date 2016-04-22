package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNetinfo(t *testing.T) {
	expected := Netinfo{up: float64(123.123), down: float64(123.123)}
	actual := getNetinfo()
	assert.IsType(t, expected, actual, "`getNetinfo()` should return a Netinfo object")
	//assert.Equal(t, expected, actual, "`getNetinfo()` should be equal to --> Netinfo{up: 123.123, down: 123.123}")
}
