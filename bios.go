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
func getBIOSinfo() (ret biosinfo, err error) { //TODO add other systems than Linux
	var err1, err2, err3, err4, err5 error

	ret.boardName, err1 = readAndTrimFile("/sys/devices/virtual/dmi/id/board_name")
	ret.boardVendor, err2 = readAndTrimFile("/sys/devices/virtual/dmi/id/board_vendor")
	err = appendError(err1, err2)
	ret.biosVendor, err3 = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_vendor")
	err = appendError(err, err3)
	ret.biosVersion, err4 = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_version")
	err = appendError(err, err4)
	ret.biosDate, err5 = readAndTrimFile("/sys/devices/virtual/dmi/id/bios_date")
	err = appendError(err, err5)

	return ret, err
}

//TODO use `sysctl` and/or kenv interfaces`on BSD
//TODO use `system_profiler` on osx
//TODO use `wmic bios get smbiosbiosversion` on windows
// We do not ask for the board serial number because it need `sudo` on Linux
// and because it's rarely usefull to look at it on a 'quick' way

//err = fmt.Errorf("%v %v %v %v %v ", err1, err2, err3, err4, err4)
