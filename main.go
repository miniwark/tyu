// `tyu` is a small command line utility to display basic system informations
// in the terminal.
package main

import (
	"strconv"

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

	// display information about the CPU usage
	cpuGauge := ui.NewGauge()
	cpuGauge.BorderLabel = "CPU usage "
	cpuGauge.BarColor = ui.ColorBlue
	cpuGauge.Width = 39
	cpuGauge.Height = 3
	cpuGauge.X = 0
	cpuGauge.Y = 6

	// display informations about the physical disks
	diskGauges := make([]*ui.Gauge, 3) // 3 disks max
	for i := range diskGauges {        //TEMP range = 3
		diskGauges[i] = ui.NewGauge()
		diskGauges[i].BarColor = ui.ColorBlue
		diskGauges[i].Width = 39
		diskGauges[i].Height = 3
		diskGauges[i].X = 0
		diskGauges[i].Y = 9 + (i * 3)
	}

	// display information about the global network activity (all interfaces)
	netinfo := ui.NewList()
	netinfo.BorderLabel = "Network "
	netitems := []string{
		"[Up:   ](fg-green)",
		"[Down: ](fg-green)",
	}
	netinfo.Items = netitems
	netinfo.Width = 39
	netinfo.Height = 4
	netinfo.X = 0
	netinfo.Y = 18

	// display system informations about the host
	host := getHostinfo()
	hostinfo := ui.NewList()
	hostinfo.BorderLabel = "Host "
	hostitems := []string{
		"[Hostname         ](fg-cyan)" + host.hostname,
		"[Domain           ](fg-cyan)" + host.domainname,
		"[OS               ](fg-cyan)" + host.os,
		"[OS version       ](fg-cyan)" + host.osRelease,
		"[Platform         ](fg-cyan)" + host.platform,
		"[Platform version ](fg-cyan)" + host.platformVersion,
		"[Architecture     ](fg-cyan)" + host.arch,
		"[Uptime           ](fg-cyan)",
	}
	hostinfo.Items = hostitems
	hostinfo.Width = 39
	hostinfo.Height = 10
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
		"[Frequency   ](fg-cyan)" + cpu.cpuMhz + " Mhz",
		"[Temperature ](fg-cyan)", //TODO
	}
	cpuinfo.Width = 39
	cpuinfo.Height = 7
	cpuinfo.X = 40
	cpuinfo.Y = 10

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
	biosinfo.Y = 17

	// display a quit help text
	quit := ui.NewPar("[Type 'q' to exit](fg-white,bg-blue)")
	quit.Height = 1
	quit.Width = 39
	quit.X = 1
	quit.Y = 23
	quit.Border = false

	// variables to calculate network trafic (to avoid cumulative display)
	var networkOld, networkNew Netinfo

	// render the dashboard with 26x80 fixed size
	draw := func(t int) {
		// update memory usage gauges
		mem := getMeminfo()
		ramGauge.Percent = mem.ramUsedPercent
		ramGauge.Label = "{{percent}}% - " + mem.ramUsed + "/" + mem.ramTotal + " GiB"
		swapGauge.Percent = mem.swapUsedPercent
		swapGauge.Label = "{{percent}}% - " + mem.swapUsed + "/" + mem.swapTotal + " GiB"

		// update CPU usage gauge
		cpuGauge.Percent = getCPUpercent()

		// update disks usage gauges
		disk := getDiskinfo()
		for i := range disk {
			if i >= 3 { // display 3 disk max
				break
			}
			diskGauges[i].BorderLabel = disk[i].device + " disk usage "
			diskGauges[i].Percent = disk[i].usedPercent
			diskGauges[i].Label = "{{percent}}% - " + disk[i].used + "/" + disk[i].total + " GiB"
		}

		// update network informations
		net := getNetinfo()
		networkNew.up = net.up - networkOld.up
		networkNew.down = net.down - networkOld.down
		networkOld = net

		netitems[0] = "[Up:   ](fg-green)" + strconv.FormatFloat(networkNew.up, 'f', 1, 64) + " KiB"
		netitems[1] = "[Down: ](fg-green)" + strconv.FormatFloat(networkNew.down, 'f', 1, 64) + " KiB"

		// update the host informations
		hostitems[7] = "[Uptime           ](fg-cyan)" + getUptime()

		// register the gauges and blocks to the renderer
		ui.Render(
			ramGauge,
			swapGauge,
			cpuGauge,
			diskGauges[0],
			diskGauges[1],
			diskGauges[2],
			netinfo,
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
