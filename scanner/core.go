package core

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func StartScan(host string) {

	sliceLength := len(LargePortList)

	concurrency := 20

	var wg sync.WaitGroup

	wg.Add(sliceLength)

	semaphore := make(chan struct{}, concurrency)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()
			val := LargePortList[i]
			getPorts(host, val)
		}(i)
	}
	wg.Wait()
}

func getPorts(host string, port int) {
	ms := 1500
	timeOut := time.Duration(ms) * time.Millisecond
	_, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeOut)

	if err != nil {
		fmt.Printf("Closed port on host %s and port %d.\n", host, port)
	} else {
		fmt.Printf("Open port on host %s and port %d.\n", host, port)
	}
}

func GetIP(Hostname string) (net.IP, error) {
	addr, err := net.LookupIP(Hostname)

	if err != nil {
		return nil, errors.New("Bad host" + err.Error())
	}

	for _, ip := range addr {
		if ip.To4() != nil {
			return ip, nil
		}
	}

	return nil, errors.New("Failure on getting IP")
}
