package main

import (
	"strconv"

	ui "github.com/gizak/termui"
)

// createHostList display system informations about the host
func createHostList() *ui.List {
	host, _ := getHostStat()

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
	l.X = 0
	l.Y = 0
	return l
}

// 	updateHostList/ update the host informations
func updateHostList(g *ui.List) {
	uptime, _ := getUptime()
	g.Items[7] = "[Uptime           ](fg-cyan)" + uptime
}

// createCPUList display informations about the CPUs
func createCPUList() *ui.List {
	cpu, _ := getCPUStat()

	l := ui.NewList()
	l.BorderLabel = "CPU "
	l.Items = []string{
		"[CPUs        ](fg-cyan)" + cpu.count,
		"[Vendor      ](fg-cyan)" + cpu.vendorID,
		"[Model       ](fg-cyan)" + cpu.modelName, //TODO use refreshing rate to display roll long text ?
		"[Frequency   ](fg-cyan)" + cpu.mhz + " Mhz",
		"[Temperature ](fg-cyan)", //TODO CPU temperature
	}
	l.Width = 39
	l.Height = 7
	l.X = 0
	l.Y = 10
	return l
}

// 	updateCPUList update the CPU informations
func updateCPUList(g *ui.List) {
	//uptime := getCPUtemp() //TODO CPU temperature
	g.Items[4] = "[Temperature ](fg-cyan)" + "--"
}

// createBIOSList display bios and motherboard informations
func createBIOSList() *ui.List {
	bios, _ := getBIOSStat()

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
	l.X = 0
	l.Y = 17
	return l
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
	l.X = 40
	l.Y = 0
	return l
}

// NetworkOld keep track of network traffic up & down
var NetworkOld NetStat //TODO get rid of this globals

//  updateNetList update network informations
func updateNetList(g *ui.List) {
	net, _ := getNetStat()
	up := net.up - NetworkOld.up
	down := net.down - NetworkOld.down
	NetworkOld = net

	g.Items[0] = "[Up   ](fg-cyan)" + strconv.FormatFloat(up, 'f', 1, 64) + " KiB"
	g.Items[1] = "[Down ](fg-cyan)" + strconv.FormatFloat(down, 'f', 1, 64) + " KiB"
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
	l.X = 60
	l.Y = 0
	return l
}

// updateProcList update processes informations
func updateProcList(g *ui.List) {
	procs, _ := getProcStat()
	g.Items[0] = "[Tasks   ](fg-cyan)" + procs.total
	g.Items[1] = "[Running ](fg-cyan)" + procs.running
}

// createRamGauge display physical RAM usage informations
func createRAMGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "RAM usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 40
	g.Y = 4
	return g
}

// updateRAMGauge update the percentage and label of the RAM gauge
func updateRAMGauge(g *ui.Gauge) {
	ram, _ := getRAMStat()

	g.Percent = ram.usedPercent
	g.Label = "{{percent}}% - " + ram.used + "/" + ram.total + " GiB"
}

// createSwapGauge display physical swap usage informations
func createSwapGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "Swap usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 40
	g.Y = 7
	return g
}

// updateSwapGauge update the percentage and label of the Swap gauge
func updateSwapGauge(g *ui.Gauge) {
	swap, _ := getSwapStat()

	g.Percent = swap.usedPercent
	g.Label = "{{percent}}% - " + swap.used + "/" + swap.total + " GiB"
}

// createCPUGauge display information about the CPU usage
func createCPUGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.BorderLabel = "CPU usage "
	g.BarColor = ui.ColorBlue
	g.Width = 39
	g.Height = 3
	g.X = 40
	g.Y = 10
	return g
}

// updateRAMGauge update the percentage of the CPU gauge
func updateCPUGauge(g *ui.Gauge) {
	g.Percent, _ = getCPUPercent()
}

// createDiskGauges display informations about the physical disks
// up to 3 disks. return an arrau of termui.Gauge
func createDiskGauge() []*ui.Gauge {
	disk, _ := getDiskStat()

	g := make([]*ui.Gauge, 3) // display  3 disks max

	for i := range disk {
		if i >= 3 { // display 3 disk max
			break
		}
		g[i] = ui.NewGauge()
		g[i].BarColor = ui.ColorBlue
		g[i].Width = 39
		g[i].Height = 3
		g[i].X = 40
		g[i].Y = 14 + (i * 3)
	}
	return g
}

// updateRAMGauge update the percentages of the disk gauges
func updateDiskGauge(g []*ui.Gauge) { //BUG tyu crash when mounting an external disk
	disk, _ := getDiskStat()

	for i := range disk {
		if i >= 3 { // display 3 disk max
			break
		}
		g[i].BorderLabel = disk[i].device + " disk usage "
		g[i].Percent = disk[i].usedPercent
		g[i].Label = "{{percent}}% - " + disk[i].used + "/" + disk[i].total + " GiB"
	}
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

// init the termui dashboard
func dashboard() {
	// init the dashboard
	if err := ui.Init(); err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	defer ui.Close()

	// inits the dashboard widgets
	hostList := createHostList()
	cpuList := createCPUList()
	biosList := createBIOSList()
	netList := createNetList()
	procList := createProcList()
	ramGauge := createRAMGauge()
	swapGauge := createSwapGauge()
	cpuGauge := createCPUGauge()
	diskGauge := createDiskGauge()
	quitPar := createQuitPar()

	// render the dashboard with 26x80 fixed size
	draw := func(t int) {
		// updates the dashboard widgets
		updateHostList(hostList)
		updateCPUList(cpuList)
		updateNetList(netList)
		updateProcList(procList)
		updateRAMGauge(ramGauge)
		updateSwapGauge(swapGauge)
		updateCPUGauge(cpuGauge)
		updateDiskGauge(diskGauge)

		// register the dashboard widgets (except the disks)
		ui.Render(
			hostList,
			cpuList,
			biosList,
			netList,
			procList,
			ramGauge,
			swapGauge,
			cpuGauge,
			quitPar,
		)

		// register the dashboard disk gauges widgets
		disk, _ := getDiskStat() // TODO try to avoid this check
		if len(disk) >= 1 {      // one disk or more
			ui.Render(diskGauge[0])
		}
		if len(disk) >= 2 { // two disks or more
			ui.Render(diskGauge[1])
		}
		if len(disk) == 3 { // three disks
			ui.Render(diskGauge[2])
		}
	}

	// setup the events handlers
	ui.Handle("/sys/kbd/q", func(ui.Event) { // quit on `q` keystroke
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-c", func(ui.Event) { // quit on `CTRL+c` keystroke
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) { // refresh every second
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})

	ui.Loop()
}
