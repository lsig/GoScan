package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	// Process the input
	fmt.Println("Ports:", portList)
	fmt.Println("IP Addresses:", ipAddresses)
}
