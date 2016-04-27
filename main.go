// tyu is a small command line utility to display basic system informations
// in the terminal.
package main

import (
	"fmt"
	"os"
	"strconv"

	ui "github.com/gizak/termui"
	flag "github.com/ogier/pflag"
)

// command line argument
var versionFlag bool

// command line arguments setup
func init() {
	flag.BoolVarP(&versionFlag, "version", "V", false, "show program's version number and exit")
	// command line help text
	help := "Usage:\n  tyu [-h] [-V]\n\n" +
		"Options:\n" +
		"  -h, --help            show this help message and exit\n" +
		"  -V, --version         show program's version number and exit\n\n"
	flag.Usage = func() { // overide the ugly default usage message
		fmt.Fprintf(os.Stderr, help)
	}
}

func main() {
	// command line arguments
	flag.Parse()
	// show program's version number and exit
	if versionFlag {
		fmt.Println("Tyu version " + tyuVersion)
		os.Exit(0)
	}

	// init termui
	if err := ui.Init(); err != nil {
		panic(err) //TODO do not panic but manage the error
	}
	defer ui.Close()

	ramGauge := createRAMGauge()
	swapGauge := createSwapGauge()
	cpuGauge := createCPUGauge()

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

	netinfo := createNetList()
	procinfo := createProcList()
	hostinfo := createHostList()
	cpuinfo := createCPUList()
	biosinfo := createBIOSList()

	quit := createQuitPar()

	// variables to calculate network traffic (to avoid cumulative display)
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
		netinfo.Items[0] = "[Up   ](fg-cyan)" + strconv.FormatFloat(networkNew.up, 'f', 1, 64) + " KiB"
		netinfo.Items[1] = "[Down ](fg-cyan)" + strconv.FormatFloat(networkNew.down, 'f', 1, 64) + " KiB"

		// update processes informations
		procs := getProcinfo()
		procinfo.Items[0] = "[Tasks   ](fg-cyan)" + procs.total
		procinfo.Items[1] = "[Running ](fg-cyan)" + procs.running

		// update the host informations
		uptime := getUptime()
		hostinfo.Items[7] = "[Uptime           ](fg-cyan)" + uptime

		// update the CPUs informations
		cpuinfo.Items[4] = "[Temperature ](fg-cyan)" + "--"

		// register the gauges and blocks to the renderer
		ui.Render(
			ramGauge,
			swapGauge,
			cpuGauge,
			diskGauges[0],
			diskGauges[1],
			diskGauges[2],
			netinfo,
			procinfo,
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
