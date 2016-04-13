package main

import (
	"io/ioutil"
	"strings"
)

// `biosinfo` represent the computer motherboard and bios informations
type biosinfo struct {
	boardName   string // name of the motherboard ex. 'C-64'
	boardVendor string // name of the motherboard  vendor ex. 'Commodore'
	biosVendor  string // name of the BIOS editor ex. 'Coreboot'
	biosVersion string // version of the BIOS
	biosDate    string // release date of the BIOS
}

// Get informations about the BIOS (or UEFI)
// we prefer to read directly from `/sys/devices/virtual/dmi/id/` than `demidecode` to avoid `sudo`
func getBIOSinfo() biosinfo {
	boardname, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	boardvendor, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_vendor")
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	biosvendor, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_vendor")
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	biosversion, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_version")
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	biosdate, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_date")
	if err != nil {
		panic(err) //TODO do not panic but manage the error
	}

	ret := biosinfo{
		boardName:   strings.TrimRight(string(boardname), "\n"), // Trimright remove the EOF carriage return
		boardVendor: strings.TrimRight(string(boardvendor), "\n"),
		biosVendor:  strings.TrimRight(string(biosvendor), "\n"),
		biosVersion: strings.TrimRight(string(biosversion), "\n"),
		biosDate:    strings.TrimRight(string(biosdate), "\n"),
	}
	return ret
}

//TODO use `sysctl` and/or kenv interfaces`on BSD
//TODO use `system_profiler` on osx
//TODO use `wmic bios get smbiosbiosversion` on windows
// We do not ask for the board serial number because it need `sudo` on Linux
// and because it's rarely usefull to look at it on a 'quick' way
