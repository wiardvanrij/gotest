package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/wiardvanrij/testing/scanner"
)

var Ports []int

type Commands struct {
	Hostname string `long:"hostname" description:"A hostname" required:"true"`
	Ports    string `short:"p"  long:"port" description:"Small, Medium, Large, Xlarge or custom comma seperated list" required:"true"`
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

	if val, ok := scanner.Ports[commands.Ports]; ok {
		Ports = val
	} else {
		for _, p := range strings.Split(commands.Ports, ",") {
			port, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				fmt.Println("Invalid port range")
				os.Exit(1)
			}
			if port < 1 || port > 65535 {
				fmt.Println("Invalid port")
				os.Exit(1)
			}
			Ports = append(Ports, port)
		}
	}

	if ip, err := scanner.GetIP(commands.Hostname); err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		scanner.StartScan(ip.String(), Ports)
	}
}
