package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the port flag
	var ports string
    flag.StringVar(&ports, "p", "", "Optional: List of ports separated by commas")

    // Parse the flags
    flag.Parse()

    var portList []string
    if ports != "" {
        // Split the ports into a slice if provided
        portList = strings.Split(ports, ",")
    }

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

