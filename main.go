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

	// display physical RAM and swap memory informations
	memGauge := ui.NewGauge() //TODO try to draw the border around gauge, used, free and total
	memGauge.BorderLabel = "Memory usage "
	memGauge.BarColor = ui.ColorBlue
	memGauge.Width = 39
	memGauge.Height = 3
	memGauge.X = 0
	memGauge.Y = 0

	memUsed := ui.NewPar("Used")
	memUsed.Height = 1
	memUsed.Width = 13
	memUsed.X = 1
	memUsed.Y = 3
	memUsed.Border = false

	memFree := ui.NewPar("Free")
	memFree.Height = 1
	memFree.Width = 13
	memFree.X = 14
	memFree.Y = 3
	memFree.Border = false

	memTotal := ui.NewPar("Total")
	memTotal.Height = 1
	memTotal.Width = 13
	memTotal.X = 27
	memTotal.Y = 3
	memTotal.Border = false

	swapGauge := ui.NewGauge()
	swapGauge.BorderLabel = "Swap usage "
	swapGauge.BarColor = ui.ColorBlue
	swapGauge.Width = 39
	swapGauge.Height = 3
	swapGauge.X = 0
	swapGauge.Y = 4

	swapUsed := ui.NewPar("Used")
	swapUsed.Height = 1
	swapUsed.Width = 13
	swapUsed.X = 1
	swapUsed.Y = 7
	swapUsed.Border = false

	swapFree := ui.NewPar("Free")
	swapFree.Height = 1
	swapFree.Width = 13
	swapFree.X = 14
	swapFree.Y = 7
	swapFree.Border = false

	swapTotal := ui.NewPar("Total")
	swapTotal.Height = 1
	swapTotal.Width = 13
	swapTotal.X = 27
	swapTotal.Y = 7
	swapTotal.Border = false

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

	// display informations about the RAM and swap memory
	cpu := getCPUinfo()
	cpuinfo := ui.NewList()
	cpuinfo.BorderLabel = "CPU "
	cpuinfo.Items = []string{
		"[CPUs   ](fg-cyan)" + cpu.count, //TODO review item names compared to other cpu utilities
		"[Vendor ](fg-cyan)" + cpu.vendorID,
		"[Model  ](fg-cyan)" + cpu.modelName,
		"[Speed  ](fg-cyan)" + cpu.cpuMhz + " Mhz",
	}
	cpuinfo.Width = 39
	cpuinfo.Height = 6
	cpuinfo.X = 40
	cpuinfo.Y = 9

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
		memGauge.Percent = mem.memUsedPercent
		memUsed.Text = "[Used](fg-cyan) " + mem.memUsed + "MB"
		memFree.Text = "[Free](fg-cyan) " + mem.memFree + "MB"
		memTotal.Text = "[Total](fg-cyan) " + mem.memTotal + "MB"
		swapGauge.Percent = mem.swapUsedPercent
		swapUsed.Text = "[Used](fg-cyan) " + mem.swapUsed + "MB"
		swapFree.Text = "[Free](fg-cyan) " + mem.swapFree + "MB"
		swapTotal.Text = "[Total](fg-cyan) " + mem.swapTotal + "MB"

		ui.Render(
			memGauge,
			memUsed,
			memFree,
			memTotal,
			swapGauge,
			swapUsed,
			swapFree,
			swapTotal,
			hostinfo,
			cpuinfo,
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
