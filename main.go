// `tyu` is a small command line utility to display basic system informations
// in the terminal.
package main

import (
	ui "github.com/gizak/termui"
)

func main() {
	//init termui
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	// quit helptext paragraph widget
	p0 := ui.NewPar("[Type 'q' to exit](fg-white,bg-blue)")
	p0.Height = 1
	p0.Width = 39
	p0.X = 1
	p0.Y = 23
	p0.Border = false

	// render the dashboard
	draw := func(t int) {
		ui.Render(p0)
	}

	// quit on `q` keystroke handler
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	// quit on `CTRL+c` keystroke handler
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	// timer handler, refresh every second
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})
	ui.Loop()
}
