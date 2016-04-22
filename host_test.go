package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostinfo(t *testing.T) {
	expected := hostinfo{
		hostname:        "abc",
		domainname:      "abc",
		os:              "abc",
		osRelease:       "abc",
		platform:        "abc",
		platformVersion: "abc",
		arch:            "abc",
	}
	actual := getHostinfo()
	assert.IsType(t, expected, actual, "`getHostinfo()` should return a `hostinfo` type")
}
