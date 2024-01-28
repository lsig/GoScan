package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/lsig/PortScanner/port"
	"github.com/lsig/PortScanner/utils"
)

// https://stackoverflow.com/questions/8509152/max-number-of-goroutines
const maxGoroutines = 10000

func main() {
	// Define the port flag
	portsPtr := flag.String("p", "", "Optional: List of ports separated by commas")
	// Parse flag
	flag.Parse()

	portList := strings.Split(*portsPtr, ",")

	// Remaining arguments are the IP addresses
	ipAddresses := flag.Args()

	if len(ipAddresses) == 0 {
		fmt.Println("No IP addresses provided.")
		os.Exit(1)
	}

	ips := utils.ConvertArgsToIPs(ipAddresses)
	ports := utils.ConvertFlagToPorts(portList)

	var wg sync.WaitGroup
	// https://levelup.gitconnected.com/go-concurrency-pattern-semaphore-9587d45f058d
	sem := make(chan struct{}, maxGoroutines)

	for _, ip := range ips {
		for _, po := range ports {
			wg.Add(1)
			// Scan port concurrently
			go func(host net.IP, portno string) {
				defer wg.Done()
				sem <- struct{}{} // aquire semaphore
				port.Scan(host, portno)
				<-sem // release semaphore

			}(ip, po)
		}
	}

	wg.Wait()
	fmt.Println("Scanning complete.")
}
