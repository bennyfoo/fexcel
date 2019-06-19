package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/onerobotics/fexcel/commenter"
)

const VERSION = "1.0-beta.6"

var (
	defaultSheet string
	offset       int
	cfg          commenter.Config
)

const logo = `  __                  _
 / _|                | |
| |_ _____  _____ ___| |
|  _/ _ \ \/ / __/ _ \ |
| ||  __/>  < (_|  __/ |
|_| \___/_/\_\___\___|_|
             v1.0-beta.6

by ONE Robotics Company
www.onerobotics.com

`

func usage() {
	fmt.Fprintf(os.Stderr, logo)
	fmt.Fprintf(os.Stderr, "Usage: fexcel [options] filename host(s)...\n\n")

	fmt.Fprintf(os.Stderr, "Example: fexcel -sheet Data -numregs A2 -posregs Sheet2:D2 -timeout 1000 spreadsheet.xlsx 127.0.0.101 127.0.0.102\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	flag.StringVar(&defaultSheet, "sheet", "Sheet1", "the name of the default sheet to look at if the sheet is not specified for an item")
	flag.IntVar(&offset, "offset", 1, "column offset from ids to comments")
	flag.StringVar(&cfg.Numregs, "numregs", "", "start cell of numeric register ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Posregs, "posregs", "", "start cell of position register ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Ualms, "ualms", "", "start cell of user alarm ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Rins, "rins", "", "start cell of robot input ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Routs, "routs", "", "start cell of robot output ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Dins, "dins", "", "start cell of digital input ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Douts, "douts", "", "start cell of digital output ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Gins, "gins", "", "start cell of group input ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Gouts, "gouts", "", "start cell of group output ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Ains, "ains", "", "start cell of analog input ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Aouts, "aouts", "", "start cell of analog output ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Sregs, "sregs", "", "start cell of string register ids (e.g. A2 or Sheet2:A2)")
	flag.StringVar(&cfg.Flags, "flags", "", "start cell of flag ids (e.g. A2 or Sheet2:A2)")
	flag.IntVar(&cfg.Timeout, "timeout", 500, "timeout value in milliseconds")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
	}

	filename := args[0]
	if filename == "" {
		usage()
	}

	ext := filepath.Ext(filename)
	if ext != ".xlsx" {
		fmt.Println("Error: fexcel only supports .xlsx files generated by Excel 2007 or later")
		os.Exit(1)
	}

	hosts := args[1:]

	fmt.Printf(logo)

	c, err := commenter.New(filename, defaultSheet, offset, cfg)
	check(err)

	for _, host := range hosts {
		err = c.Update(host)
		check(err)

		fmt.Println()
	}
}
