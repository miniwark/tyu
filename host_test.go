package main

import (
	"syscall"
	"testing"

	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/assert"
)

// TestGetHostStat test the returned fields values of `getHostStat()`
func TestGetHostStat(t *testing.T) {
	// setup the faking of `host.Info()`
	oldHostInfo := hostInfo
	hostInfo = func() (*host.InfoStat, error) {
		ret := &host.InfoStat{
			Hostname:        "abc",
			OS:              "linux", // we need a `Linux` OS if we want to test datas retrivied by `getUname()`
			Platform:        "abc",
			PlatformVersion: "abc",
		}
		return ret, nil
	}
	oldgetUname := getUname
	getUname = func() (syscall.Utsname, error) {
		int8string := [65]int8{ // See the ASCII table: 97 = "a" ; 98 ="b" ; 99 = "c"
			64, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
			64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
			64, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
		} //TODO if possible avoid this complicate way to test because of [65]int8 requirement by syscall.Utsname
		ret := syscall.Utsname{
			Release:    int8string,
			Machine:    int8string,
			Domainname: int8string,
		}
		return ret, nil
	}

	// test
	expected := hostStat{
		hostname:        "abc",
		os:              "Linux",
		platform:        "Abc",
		platformVersion: "abc",
		domainname:      "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
		osRelease:       "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
		arch:            "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
	}
	actual, err := getHostStat()

	assert.NoError(t, err, "`getHostStat()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getHostStat()` should be equal to main.hostStat{hostname:\"hostname\", domainname:\"domainname\", os:\"Os\", osRelease:\"1.0.0\", platform:\"Platform\", platformVersion:\"10.00\", arch:\"x86_64\"}")

	// teardown
	hostInfo = oldHostInfo
	getUname = oldgetUname
}

// TestGetHostStatType test if `getHostStat()` return a `hostStat` type and if each fields have the correct types
func TestGetHostStatType(t *testing.T) {
	expected := hostStat{
		hostname:        "", // the result values of the fields are not tested
		domainname:      "",
		os:              "",
		osRelease:       "",
		platform:        "",
		platformVersion: "",
		arch:            "",
	}
	actual, _ := getHostStat()

	assert.IsType(t, expected, actual, "`getHostStat()` should return a `hostStat` type")
	assert.IsType(t, expected.hostname, actual.hostname, "`getHostStat()` should return a `hostname` field with a string type")
	assert.IsType(t, expected.domainname, actual.domainname, "`getHostStat()` should return a `domainname` field with a string type")
	assert.IsType(t, expected.os, actual.os, "`getHostStat()` should return a `os` field with a string type")
	assert.IsType(t, expected.osRelease, actual.osRelease, "`getHostStat()` should return a `osRelease` field with a string type")
	assert.IsType(t, expected.platform, actual.platform, "`getHostStat()` should return a `platform` field with a string type")
	assert.IsType(t, expected.platformVersion, actual.platformVersion, "`getHostStat()` should return a `platformVersion` field with a string type")
	assert.IsType(t, expected.arch, actual.arch, "`getHostStat()` should return a `arch` field with a string type")
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
	actual, err := getUptime()

	assert.NoError(t, err, "`getUptime()` should not have returned an error")
	assert.Equal(t, expected, actual, "`getUptime` should be equal to \"24h0m0s\"")

	// teardown
	hostUptime = oldHostUptime
}

// TestGetUptimeType test if `getUptime()` return a value with a string type
func TestGetUptimeType(t *testing.T) {
	expected := "" // the result value is not tested
	actual, _ := getUptime()

	assert.IsType(t, expected, actual, "`getUptime` should return a string")
}

// TestGetUnameType test if `getUname()` return a value with a syscall.Utsname type
func TestGetUname(t *testing.T) {
	expected := syscall.Utsname{}
	actual, _ := getUname() //TODO add test for error case

	assert.IsType(t, expected, actual, "`getUname()` should return a syscall.Utsname type")
}

// TestHostInfo test if `hostInfo()` return a value with a *host.InfoStat type
func TestHostInfo(t *testing.T) {
	expected := &host.InfoStat{}
	actual, _ := hostInfo()

	assert.IsType(t, expected, actual, "`hostInfo()` should return a *host.InfoStat type")
}

// TestHhostUptime test if `hostUptime()` return a value with a uint64 type
func TestHostUptime(t *testing.T) {
	expected := uint64(0) // the result value is not tested
	actual, _ := hostUptime()

	assert.IsType(t, expected, actual, "`hostUptime()` should return a uint64 type")
}
