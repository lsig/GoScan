package port

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func ScanPort(ipAddr net.IP, port string) {
	address := net.JoinHostPort(ipAddr.String(), port)
	dialer := net.Dialer{Timeout: 2 * time.Second} // Set a timeout
	conn, err := dialer.Dial("tcp", address)

	if err == nil {
		fmt.Printf("%s open\n", address)

		if err != nil {
			fmt.Println("Error setting timeout", err)
		}

		_, err = conn.Write([]byte("71"))

		if err != nil {
			fmt.Printf("Port %s did not allowe the probe\n", port)
		}

		conn.Close()
	} else if strings.Contains(err.Error(), "connection refused") {
		fmt.Printf("%s closed\n", address)
	} else {
		fmt.Printf("%s", err.Error())
	}
}
