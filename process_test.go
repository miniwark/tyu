package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProcinfo(t *testing.T) {
	expected := procinfo{
		total:   string(100),
		running: string(100),
	}
	actual := getProcinfo()
	assert.IsType(t, expected, actual, "`getProcinfo` should return a `procinfo` type")
}
