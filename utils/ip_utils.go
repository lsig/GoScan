package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
)

// Converts Arguments (hostnames, IP addresses and CIDR subnets) to a list of valid IP address
//
// If args contains an invalid hostname, IP or subnet it will be omitted and an error will be printed to stderr
func ConvertArgsToIPs(args []string) []net.IP {
	ips := []net.IP{}

	for _, ip := range args {
		switch {
		case isValidIPv4(ip):
			newIPs, _ := convertToIP(ip)
			ips = append(ips, newIPs...)
		case isValidCIDR(ip):
			newIPs, _ := convertSubnetToIPs(ip)
			ips = append(ips, newIPs...)
		case isValidHostname(ip):
			newIPs, err := resolveHost(ip)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				continue
			}

			ips = append(ips, newIPs...)
		default:
			err := fmt.Errorf("IP or Hostname incorrectly formatted: %v", ip)
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
	return ips
}

// Converts hostname to the related IPv4 addresses of the host
//
// # Example
//
// "mbl.is" converts to []net.IP{"92.43.192.120"}
func resolveHost(hostname string) ([]net.IP, error) {
	ips, err := net.LookupIP(hostname)

	if err != nil {
		return nil, errors.New("Could not resolve host")
	}

	var ipv4Addresses []net.IP
	for _, ip := range ips {
		if ipV4 := ip.To4(); ipV4 != nil {
			ipv4Addresses = append(ipv4Addresses, ipV4)
		}
	}

	return ipv4Addresses, nil
}

// Converts an IP string to a valid IPv4 address
func convertToIP(ipStr string) ([]net.IP, error) {
	ip := net.ParseIP(ipStr).To4()

	if ip == nil {
		return nil, errors.New("Invalid IpV4 Address")
	}

	return []net.IP{ip}, nil
}

// Converts a CIDR to a slice of valid IPv4 addresses in the subnet
func convertSubnetToIPs(subnet string) ([]net.IP, error) {
	ip, ipnet, err := net.ParseCIDR(subnet)

	if err != nil {
		return nil, errors.New("Invalid subnet")
	}

	if ip.To4() == nil {
		return nil, errors.New("Not a valid IPv4 subnet")
	}

	var ips []net.IP

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ip = nextIP(ip, 1) {
		if ipV4 := ip.To4(); ipV4 != nil {
			ips = append(ips, ipV4)
		}
	}

	// Remove network and broadcast address
	ips = ips[1 : len(ips)-1]

	return ips, nil

}

// Finds the next IPv4 after a specified IPv4 address
//
// https://stackoverflow.com/questions/31191313/how-to-get-the-next-ip-address
func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}

// Checks if a hostname is of a valid format
func isValidHostname(hostname string) bool {
	// Pattern for matching a valid hostname (RFC 1123)
	// ChatGPT suggested this and tests in ip_utils_test.go confirm functionality
	hostnamePattern := regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.?([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])$`)
	return hostnamePattern.MatchString(hostname) && len(hostname) <= 253
}

// Checks if a string represents a valid IPv4 address
func isValidIPv4(ip string) bool {
	ipv4 := net.ParseIP(ip).To4()
	return ipv4 != nil
}

// Checks if a string represents a valid IPv4 subnet
func isValidCIDR(cidr string) bool {
	ip, _, err := net.ParseCIDR(cidr)
	ipv4 := ip.To4()
	return err == nil && ipv4 != nil
}
