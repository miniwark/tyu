package main

// `biosinfo` represent the computer informations
type biosinfo struct {
	motherboard  string
	productName  string
	serialNumber string //TODO will probably not be available without root
	BIOS         string
	BIOSVersion  string
}

// Get informations about the BIOS (or UEFI)
func getBIOSinfo() biosinfo {
	result := biosinfo{
		motherboard:  "",
		productName:  "",
		serialNumber: "",
		BIOS:         "",
		BIOSVersion:  "",
	}
	return result
}

//TODO read directly from   /sys/devices/virtual/dmi/id/* /sys/class/dmi/id/* on linux to avoid demidecode sudo problem
//or try lshal see http://stackoverflow.com/questions/20206474/extract-the-linux-serial-number-without-sudo
//or dmesg ?
//TODO use `sysctl` and/or kenv interfaces`on BSD
//TODO use `system_profiler` on osx
//TODO use `wmic bios get smbiosbiosversion` on windows
