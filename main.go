package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
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

	if ip, err := getIP(commands.Hostname); err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {

		//getPorts(ip.String(), commands.Port)
		StartScan(ip.String())
	}
}

func StartScan(Host string) {

	port := make(chan int64)

	port <- 22
	go getPorts(Host, port)

	

	close(port)

}

func getPorts(host string, port chan int64) {
	ms := 500
	timeOut := time.Duration(ms) * time.Millisecond
	_, err := net.DialTimeout("tcp", host+":"+strconv.FormatInt(<-port, 10), timeOut)

	if err != nil {
		fmt.Printf("Closed port on host %s and port %d.\n", host, <-port)
		return
	} else {
		fmt.Printf("Open port on host %s and port %d.\n", host, <-port)

	}
}


func getIP(Hostname string) (net.IP, error) {
	addr, err := net.LookupIP(Hostname)

	if err != nil {
		return nil, errors.New("Bad host" + err.Error())
	}

	for _, ip := range addr {
		if ip.To4() != nil {
			return ip, nil
		}
	}

	return nil, errors.New("math: square root of negative number")

}

