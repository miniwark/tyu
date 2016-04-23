package main

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

	ret.boardName = readAndTrimFile("/sys/devices/virtual/dmi/id/board_name")
	ret.boardVendor = readAndTrimFile("/sys/devices/virtual/dmi/id/board_vendor")
	ret.biosVendor = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_vendor")
	ret.biosVersion = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_version")
	ret.biosDate = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_date")

	return ret
}

//TODO use `sysctl` and/or kenv interfaces`on BSD
//TODO use `system_profiler` on osx
//TODO use `wmic bios get smbiosbiosversion` on windows
// We do not ask for the board serial number because it need `sudo` on Linux
// and because it's rarely usefull to look at it on a 'quick' way
