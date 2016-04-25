package main

import (
	"syscall"
	"testing"

	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/assert"
)

// TODO finish this test
// TestGetMeminfo test the returned fields values of `getHostinfo()`
func TestGetHostinfo(t *testing.T) {
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
	expected := hostinfo{
		hostname:        "abc",
		os:              "Linux",
		platform:        "Abc",
		platformVersion: "abc",
		domainname:      "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
		osRelease:       "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
		arch:            "@0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZ@abcdefghijklmnopqrstuvwxyz", //
	}
	actual := getHostinfo()

	assert.Equal(t, expected, actual, "`getHostinfo()` should be equal to main.hostinfo{hostname:\"hostname\", domainname:\"domainname\", os:\"Os\", osRelease:\"1.0.0\", platform:\"Platform\", platformVersion:\"10.00\", arch:\"x86_64\"}")

	// teardown
	hostInfo = oldHostInfo
	getUname = oldgetUname
}

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

	assert.Equal(t, expected, actual, "`getUptime` should be equal to \"24h0m0s\"")

	// teardown
	hostUptime = oldHostUptime
}

// TestGetUptimeType test if `getUptime()` return a value with a string type
func TestGetUptimeType(t *testing.T) {
	expected := "" // the result value is not tested
	actual := getUptime()

	assert.IsType(t, expected, actual, "`getUptime` should return a string")
}

// TestGetUnameType test if `getUname()` return a value with a syscall.Utsname type
func TestGetUname(t *testing.T) {
	expected := syscall.Utsname{}
	actual, _ := getUname() //TODO add test for error case

	assert.IsType(t, expected, actual, "`getUname` should return a syscall.Utsname type")
}
