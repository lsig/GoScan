package port

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Scans TCP Port of a specified host
//
// If no packet is sent from the host within 2 seconds the connection times out
func Scan(ipAddr net.IP, port string) {
	address := net.JoinHostPort(ipAddr.String(), port)
	dialer := net.Dialer{Timeout: 2 * time.Second} // Set a timeout
	conn, err := dialer.Dial("tcp", address)

	if err == nil {
		fmt.Printf("%s open\n", address)
		conn.Close()
	} else if strings.Contains(err.Error(), "connection refused") {
		fmt.Printf("%s closed\n", address)
	}
}
