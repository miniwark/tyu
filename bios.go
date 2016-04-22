package main

import (
	"io/ioutil"
	"strings"
)

// biosinfo represent the computer motherboard and bios informations
type biosinfo struct {
	boardName   string // name of the motherboard ex. 'C-64'
	boardVendor string // name of the motherboard  vendor ex. 'Commodore'
	biosVendor  string // name of the BIOS editor ex. 'Coreboot'
	biosVersion string // version of the BIOS
	biosDate    string // release date of the BIOS
}

// Get informations about the BIOS (or UEFI)
// we prefer to read directly from `/sys/devices/virtual/dmi/id/` than `demidecode` to avoid `sudo`
func getBIOSinfo() biosinfo { //TODO add other systems than Linux
	ret := biosinfo{}

	boardname, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err == nil {
		ret.boardName = strings.TrimRight(string(boardname), "\n") // Trimright remove the EOF carriage return
	} else {
		ret.boardName = ""
	}

	boardvendor, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_vendor")
	if err == nil {
		ret.boardVendor = strings.TrimRight(string(boardvendor), "\n")
	} else {
		ret.boardVendor = ""
	}

	biosvendor, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_vendor")
	if err == nil {
		ret.biosVendor = strings.TrimRight(string(biosvendor), "\n")
	} else {
		ret.biosVendor = ""
	}

	biosversion, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_version")
	if err == nil {
		ret.biosVersion = strings.TrimRight(string(biosversion), "\n")
	} else {
		ret.biosVersion = ""
	}

	biosdate, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/bios_date")
	if err == nil {
		ret.biosDate = strings.TrimRight(string(biosdate), "\n")
	} else {
		ret.biosDate = ""
	}

	return ret
}

//TODO use `sysctl` and/or kenv interfaces`on BSD
//TODO use `system_profiler` on osx
//TODO use `wmic bios get smbiosbiosversion` on windows
// We do not ask for the board serial number because it need `sudo` on Linux
// and because it's rarely usefull to look at it on a 'quick' way
