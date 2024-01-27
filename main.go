package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lsig/PortScanner/port"
	"github.com/lsig/PortScanner/utils"
)

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

	for _, ip := range ips {
		for _, po := range ports {
			port.ScanPort(ip, po)
		}
	}
}
