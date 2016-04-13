// `tyu` is a small command line utility to display basic system informations
// in the terminal.
package main

import (
	ui "github.com/gizak/termui"
)

func main() {
	//init termui
	if err := ui.Init(); err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	defer ui.Close()

	// display physical RAM usage informations
	ramGauge := ui.NewGauge() //TODO try to draw the border around gauge, used, free and total
	ramGauge.BorderLabel = "RAM usage "
	ramGauge.BarColor = ui.ColorBlue
	ramGauge.Width = 39
	ramGauge.Height = 3
	ramGauge.X = 0
	ramGauge.Y = 0

	// display physical swap usage informations
	swapGauge := ui.NewGauge()
	swapGauge.BorderLabel = "Swap usage "
	swapGauge.BarColor = ui.ColorBlue
	swapGauge.Width = 39
	swapGauge.Height = 3
	swapGauge.X = 0
	swapGauge.Y = 3

	cpuGauge := ui.NewGauge()
	cpuGauge.BorderLabel = "CPU usage "
	cpuGauge.BarColor = ui.ColorBlue
	cpuGauge.Width = 39
	cpuGauge.Height = 3
	cpuGauge.X = 0
	cpuGauge.Y = 6

	// display informations about the physical disks
	// TODO only two or 3 physical disk add check and a for loop
	disk := getDiskinfo()
	disk1Gauge := ui.NewGauge()
	disk1Gauge.BorderLabel = disk[0].device + " disk usage "
	disk1Gauge.BarColor = ui.ColorBlue
	disk1Gauge.Width = 39
	disk1Gauge.Height = 3
	disk1Gauge.X = 0
	disk1Gauge.Y = 9
	disk1Gauge.Percent = disk[0].usedPercent

	disk2Gauge := ui.NewGauge()
	disk2Gauge.BorderLabel = disk[1].device + " disk usage " //TODO this fail if only one disk
	disk2Gauge.BarColor = ui.ColorBlue
	disk2Gauge.Width = 39
	disk2Gauge.Height = 3
	disk2Gauge.X = 0
	disk2Gauge.Y = 12
	disk2Gauge.Percent = disk[1].usedPercent

	// display system informations about the host
	host := getHostinfo()
	hostinfo := ui.NewList()
	hostinfo.BorderLabel = "Host "
	hostinfo.Items = []string{
		"[Hostname         ](fg-cyan)" + host.hostname,
		"[Domain           ](fg-cyan)" + host.domainname,
		"[OS               ](fg-cyan)" + host.os,
		"[OS version       ](fg-cyan)" + host.osRelease,
		"[Platform         ](fg-cyan)" + host.platform,
		"[Platform version ](fg-cyan)" + host.platformVersion,
		"[Architecture     ](fg-cyan)" + host.arch,
	}
	hostinfo.Width = 39
	hostinfo.Height = 9
	hostinfo.X = 40
	hostinfo.Y = 0

	// display informations about the CPUs
	cpu := getCPUinfo()
	cpuinfo := ui.NewList()
	cpuinfo.BorderLabel = "CPU "
	cpuinfo.Items = []string{
		"[CPUs        ](fg-cyan)" + cpu.count, //TODO review item names compared to other cpu utilities
		"[Vendor      ](fg-cyan)" + cpu.vendorID,
		"[Model       ](fg-cyan)" + cpu.modelName, //TODO use refreshing rate to display roll long text ?
		"[Frenquency  ](fg-cyan)" + cpu.cpuMhz + " Mhz",
		"[Temperature ](fg-cyan)", //TODO
	}
	cpuinfo.Width = 39
	cpuinfo.Height = 7
	cpuinfo.X = 40
	cpuinfo.Y = 9

	// display bios and motherboard informations
	bios := getBIOSinfo()
	biosinfo := ui.NewList()
	biosinfo.BorderLabel = "BIOS "
	biosinfo.Items = []string{
		"[Motherboard ](fg-cyan)" + bios.boardName,
		"[Vendor      ](fg-cyan)" + bios.boardVendor,
		"[BIOS        ](fg-cyan)" + bios.biosVendor,
		"[Version     ](fg-cyan)" + bios.biosVersion + "  " + bios.biosDate,
	}
	biosinfo.Width = 39
	biosinfo.Height = 6
	biosinfo.X = 40
	biosinfo.Y = 16

	// display a quit help text
	quit := ui.NewPar("[Type 'q' to exit](fg-white,bg-blue)")
	quit.Height = 1
	quit.Width = 39
	quit.X = 1
	quit.Y = 23
	quit.Border = false

	// render the dashboard with 26x80 fixed size
	draw := func(t int) {
		// update memory informations
		mem := getMeminfo()
		ramGauge.Percent = mem.ramUsedPercent
		ramGauge.Label = "{{percent}}% - " + mem.ramUsed + "/" + mem.ramTotal + " GiB"
		swapGauge.Percent = mem.swapUsedPercent
		swapGauge.Label = "{{percent}}% - " + mem.swapUsed + "/" + mem.swapTotal + " GiB"

		cpuGauge.Percent = getCPUpercent()

		ui.Render(
			ramGauge,
			swapGauge,
			cpuGauge,
			disk1Gauge, //TODO rename and or stack gauges together
			disk2Gauge, //TODO rename
			hostinfo,
			cpuinfo,
			biosinfo,
			quit,
		)
	}

	// quit on `q` keystroke handler
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	// quit on `CTRL+c` keystroke handler
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	// timer handler to refresh every second
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})
	ui.Loop()
}

// the dashboard is 26x80
