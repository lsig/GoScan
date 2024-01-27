package utils

import (
	"net"
	"reflect"
	"sort"
	"testing"
)

func TestConvertArgsToIPs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []net.IP
	}{
		{
			name:     "Single IPv4",
			args:     []string{"192.0.2.1"},
			expected: []net.IP{net.ParseIP("192.0.2.1").To4()},
		},
		{
			name: "Test",
			args: []string{"192.168.1.248/29", "1.1.1.1", "mbl.is", "ble124545.is"},
			expected: []net.IP{
				net.ParseIP("192.168.1.249").To4(),
				net.ParseIP("192.168.1.250").To4(),
				net.ParseIP("192.168.1.251").To4(),
				net.ParseIP("192.168.1.252").To4(),
				net.ParseIP("192.168.1.253").To4(),
				net.ParseIP("192.168.1.254").To4(),
				net.ParseIP("1.1.1.1").To4(),
				net.ParseIP("92.43.192.120").To4(),
			},
		},
		{
			name:     "Empty",
			args:     []string{""},
			expected: []net.IP{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ConvertArgsToIPs(test.args)

			sort.Slice(test.expected, func(i, j int) bool {
				return test.expected[i].String() < test.expected[j].String()
			})
			sort.Slice(actual, func(i, j int) bool {
				return actual[i].String() < actual[j].String()
			})

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("ConvertArgsToIPs(%v) = %v, want %v", test.args, actual, test.expected)
			}
		})
	}
}

func TestResolveHost(t *testing.T) {
	tests := map[string]struct {
		hostname string
		expected []net.IP
	}{
		"mbl.is": {
			hostname: "mbl.is",
			expected: []net.IP{net.ParseIP("92.43.192.120").To4()},
		},
		"visir.is": {
			hostname: "visir.is",
			expected: []net.IP{
				net.ParseIP("185.21.17.248").To4(),
				net.ParseIP("185.21.16.109").To4(),
				net.ParseIP("185.21.17.244").To4(),
				net.ParseIP("185.21.16.110").To4(),
				net.ParseIP("185.21.17.247").To4(),
				net.ParseIP("185.21.17.249").To4(),
			},
		},
		"ruv.is": {
			hostname: "ruv.is",
			expected: []net.IP{
				net.ParseIP("172.67.13.226").To4(),
				net.ParseIP("104.22.74.251").To4(),
				net.ParseIP("104.22.75.251").To4(),
			},
		},
		"ble": {
			hostname: "ble",
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, _ := resolveHost(test.hostname)

			// Sort both slices by IP string representation
			sort.Slice(test.expected, func(i, j int) bool {
				return test.expected[i].String() < test.expected[j].String()
			})
			sort.Slice(actual, func(i, j int) bool {
				return actual[i].String() < actual[j].String()
			})

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("IPs do not match: expected %v, got %v", test.expected, actual)
				t.Errorf("IPs do not match: expected %v, got %v", reflect.TypeOf(test.expected), reflect.TypeOf(actual))
			}
		})
	}
}

func TestConvertToIp(t *testing.T) {
	tests := map[string]struct {
		ipStr    string
		expected []net.IP
	}{
		"1.1.1.1": {
			ipStr:    "1.1.1.1",
			expected: []net.IP{net.ParseIP("1.1.1.1").To4()},
		},
		"ble": {
			ipStr:    "ble",
			expected: nil,
		},
		"2606:4700:10::6816:4afb": {
			ipStr:    "2606:4700:10::6816:4afb",
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, _ := convertToIP(test.ipStr)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("IP conversion does not match: expected %v, got %v", test.expected, actual)
				t.Errorf("IP conversion does not match: expected %v, got %v", reflect.TypeOf(test.expected), reflect.TypeOf(actual))
			}
		})
	}

}

func TestConvertSubnetToIPs(t *testing.T) {
	tests := map[string]struct {
		subnet   string
		expected []net.IP
	}{
		"192.168.1.248/29": {
			subnet: "192.168.1.248/29",
			expected: []net.IP{
				net.ParseIP("192.168.1.249").To4(),
				net.ParseIP("192.168.1.250").To4(),
				net.ParseIP("192.168.1.251").To4(),
				net.ParseIP("192.168.1.252").To4(),
				net.ParseIP("192.168.1.253").To4(),
				net.ParseIP("192.168.1.254").To4(),
			},
		},
		"2001:db8::/125": {
			subnet:   "2001:db8::/125",
			expected: nil,
		},
		"ble.ble.ble/125": {
			subnet:   "ble.ble.ble/125",
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, _ := convertSubnetToIPs(test.subnet)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("IP conversion does not match: expected %v, got %v", test.expected, actual)
				t.Errorf("IP conversion does not match: expected %v, got %v", reflect.TypeOf(test.expected), reflect.TypeOf(actual))
			}
		})
	}

}

func TestIsValidHostname(t *testing.T) {
	tests := map[string]struct {
		hostname string
		expected bool
	}{
		"ruv.is": {
			hostname: "ruv.is",
			expected: true,
		},
		"y8.com": {
			hostname: "y8.com",
			expected: true,
		},
		"1.1.1.1": {
			hostname: "1.1.1.1",
			expected: false,
		},
		"2606:4700:10::ac43:de2": {
			hostname: "2606:4700:10::ac43:de2",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := isValidHostname(test.hostname)

			if actual != test.expected {
				t.Errorf("Validation does not work: expected %v, got %v", test.expected, actual)
				t.Errorf("IP conversion does not match: expected %v, got host %v", test.expected, test.hostname)
			}
		})
	}
}

func TestIsValidIPv4(t *testing.T) {
	tests := map[string]struct {
		ip       string
		expected bool
	}{
		"1.1.1.1": {
			ip:       "1.1.1.1",
			expected: true,
		},
		"y8.com": {
			ip:       "y8.com",
			expected: false,
		},
		"2606:4700:10::ac43:de2": {
			ip:       "2606:4700:10::ac43:de2",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := isValidIPv4(test.ip)

			if actual != test.expected {
				t.Errorf("Validation does not work: expected %v, got %v", test.expected, actual)
				t.Errorf("IP conversion does not match: expected %v, got ip %v", test.expected, test.ip)
			}
		})
	}
}

func TestIsValidCIDR(t *testing.T) {
	tests := map[string]struct {
		cidr     string
		expected bool
	}{
		"1.1.1.1/24": {
			cidr:     "1.1.1.1/24",
			expected: true,
		},
		"1.1.1.1": {
			cidr:     "1.1.1.1",
			expected: false,
		},
		"y8.com": {
			cidr:     "y8.com",
			expected: false,
		},
		"2606:4700:10::ac43:de2/28": {
			cidr:     "2606:4700:10::ac43:de2/28",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := isValidCIDR(test.cidr)

			if actual != test.expected {
				t.Errorf("Validation does not work: expected %v, got %v", test.expected, actual)
				t.Errorf("IP conversion does not match: expected %v, got CIDR %v", test.expected, test.cidr)
			}
		})
	}
}
