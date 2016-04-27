// tyu is a small command line utility to display basic system informations
// in the terminal.
package main

import (
	"fmt"
	"os"

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

	// init and display the dashboard
	dashboard()
}
