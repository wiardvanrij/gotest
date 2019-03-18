package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	core "github.com/wiardvanrij/testing/scanner"
)

type Commands struct {
	Hostname string `long:"hostname" description:"A hostname" required:"true"`
	Port     int64  `short:"p" long:"port" description:"A port" required:"true"`
}

func main() {
	commands := Commands{}
	parser := flags.NewParser(&commands, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if ip, err := core.GetIP(commands.Hostname); err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		core.StartScan(ip.String())
	}
}
