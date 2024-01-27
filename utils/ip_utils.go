package utils

import (
	"errors"
	"net"
)

func ResolveHost(hostname string) ([]net.IP, error) {
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

func ConvertToIP(ipStr string) (net.IP, error) {
	ip := net.ParseIP(ipStr).To4()

	if ip == nil {
		return nil, errors.New("Invalid IpV4 Address")
	}

	return ip, nil
}

func ConvertSubnetToIPs(subnet string) ([]net.IP, error) {
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
