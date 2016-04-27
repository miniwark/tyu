package main

import ui "github.com/gizak/termui"

// https://github.com/divan/expvarmon/blob/master/ui_single.go
// https://github.com/mhoc/river/blob/master/src/river/send_box.go

// createRamGauge display physical RAM usage informations
func createRAMGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "RAM usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 0
	g.Y = 0
	return g
}

// createSwapGauge display physical swap usage informations
func createSwapGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "Swap usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 0
	g.Y = 3
	return g
}

// createCPUGauge display information about the CPU usage
func createCPUGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "CPU usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 0
	g.Y = 6
	return g
}

// createNetList display information about the global network activity (all interfaces)
func createNetList() *ui.List { //TODO move over ramGauge
	l := ui.NewList()
	l.BorderLabel = "Network "
	l.Items = []string{
		"[Up   ](fg-cyan)",
		"[Down ](fg-cyan)",
	}
	l.Width = 19
	l.Height = 4
	l.X = 0
	l.Y = 18
	return l
}

// createProcList display information about the processes
func createProcList() *ui.List { //TODO move over ramGauge
	l := ui.NewList()
	l.BorderLabel = "Processes "
	l.Items = []string{
		"[Tasks   ](fg-cyan)",
		"[Running ](fg-cyan)",
	}
	l.Width = 19
	l.Height = 4
	l.X = 20
	l.Y = 18
	return l
}

// createHostList display system informations about the host
func createHostList() *ui.List {
	host := getHostinfo()

	l := ui.NewList()
	l.BorderLabel = "Host "
	l.Items = []string{
		"[Hostname         ](fg-cyan)" + host.hostname,
		"[Domain           ](fg-cyan)" + host.domainname,
		"[OS               ](fg-cyan)" + host.os,
		"[OS version       ](fg-cyan)" + host.osRelease,
		"[Platform         ](fg-cyan)" + host.platform,
		"[Platform version ](fg-cyan)" + host.platformVersion,
		"[Architecture     ](fg-cyan)" + host.arch,
		"[Uptime           ](fg-cyan)",
	}
	l.Width = 39
	l.Height = 10
	l.X = 40
	l.Y = 0
	return l
}

// createCPUList display informations about the CPUs
func createCPUList() *ui.List {
	cpu := getCPUinfo()

	l := ui.NewList()
	l.BorderLabel = "CPU "
	l.Items = []string{
		"[CPUs        ](fg-cyan)" + cpu.count, //TODO review item names compared to other cpu utilities
		"[Vendor      ](fg-cyan)" + cpu.vendorID,
		"[Model       ](fg-cyan)" + cpu.modelName, //TODO use refreshing rate to display roll long text ?
		"[Frequency   ](fg-cyan)" + cpu.cpuMhz + " Mhz",
		"[Temperature ](fg-cyan)", //TODO
	}
	l.Width = 39
	l.Height = 7
	l.X = 40
	l.Y = 10
	return l
}

// createBIOSList display bios and motherboard informations
func createBIOSList() *ui.List {
	bios := getBIOSinfo()

	l := ui.NewList()
	l.BorderLabel = "BIOS "
	l.Items = []string{
		"[Motherboard ](fg-cyan)" + bios.boardName,
		"[Vendor      ](fg-cyan)" + bios.boardVendor,
		"[BIOS        ](fg-cyan)" + bios.biosVendor,
		"[Version     ](fg-cyan)" + bios.biosVersion + "  " + bios.biosDate,
	}
	l.Width = 39
	l.Height = 6
	l.X = 40
	l.Y = 17
	return l
}

// createQuitPar display a quit help text
func createQuitPar() *ui.Par {
	p := ui.NewPar("[ Type 'q' to exit ](fg-white,bg-blue)")
	p.Height = 1
	p.Width = 20
	p.X = 0
	p.Y = 23
	p.Border = false
	return p
}
