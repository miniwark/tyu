package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO finish this test
// TestGetMeminfo test the returned fields values of `getHostinfo()`
// func TestGetHostinfo(t *testing.T) {
// 	// setup the faking of `host.Info()`
// 	oldHostInfo := hostInfo
// 	hostInfo = func() (*host.InfoStat, error) {
// 		ret := &host.InfoStat{
// 			Hostname:        "hostname",
// 			OS:              "Os",
// 			Platform:        "Platform",
// 			PlatformVersion: "10.00",
// 		}
// 		return ret, nil
// 	}
//
// 	// test
// 	expected := hostinfo{
// 		hostname:        "hostname",
// 		os:              "Os",
// 		platform:        "Platform",
// 		platformVersion: "10.00",
// 		domainname:      "domainname",
// 		osRelease:       "1.0.0",
// 		arch:            "x86_64",
// 	}
// 	actual := getHostinfo()
//
// 	assert.Equal(t, expected, actual, "`getHostinfo()` should be equal to --> hostinfo{hostname:\"hostname\", domainname:\"domainname\", os:\"Os\", osRelease:\"1.0.0\", platform:\"Platform\", platformVersion:\"10.00\", arch:\"x86_64\"}")
//
// 	// teardown
// 	hostInfo = oldHostInfo
// }

// TestGetHostinfoType test if `getHostinfo()` return a `procinfo` type and if each fields have the correct types
func TestGetHostinfoType(t *testing.T) {
	expected := hostinfo{
		hostname:        "", // the result values of the fields are not tested
		domainname:      "",
		os:              "",
		osRelease:       "",
		platform:        "",
		platformVersion: "",
		arch:            "",
	}
	actual := getHostinfo()

	assert.IsType(t, expected, actual, "`getHostinfo()` should return a `hostinfo` type")
	assert.IsType(t, expected.hostname, actual.hostname, "`getHostinfo()` should return a `hostname` field with a string type")
	assert.IsType(t, expected.domainname, actual.domainname, "`getHostinfo()` should return a `domainname` field with a string type")
	assert.IsType(t, expected.os, actual.os, "`getHostinfo()` should return a `os` field with a string type")
	assert.IsType(t, expected.osRelease, actual.osRelease, "`getHostinfo()` should return a `osRelease` field with a string type")
	assert.IsType(t, expected.platform, actual.platform, "`getHostinfo()` should return a `platform` field with a string type")
	assert.IsType(t, expected.platformVersion, actual.platformVersion, "`getHostinfo()` should return a `platformVersion` field with a string type")
	assert.IsType(t, expected.arch, actual.arch, "`getHostinfo()` should return a `arch` field with a string type")
}

// TestGetUptime test the returned value of `getUptime()`
func TestGetUptime(t *testing.T) {
	// setup the faking of `cpu.Percent()`
	oldHostUptime := hostUptime
	hostUptime = func() (uint64, error) {
		ret := uint64(86400) // time.Duration to string conversion is implicitly tested --> 24h * 60m * 60s = 86400
		return ret, nil
	}

	// test
	expected := "24h0m0s"
	actual := getUptime()

	assert.Equal(t, expected, actual, "`getUptime` should be equal to --> \"24h0m0s\"")

	// teardown
	hostUptime = oldHostUptime
}

// TestGetUptimeType test if `getUptime` return a value with a string type
func TestGetUptimeType(t *testing.T) {
	expected := "" // the result value is not tested
	actual := getUptime()

	assert.IsType(t, expected, actual, "`getUptime` should return a string`")
}
