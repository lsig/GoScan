package utils

import (
	"reflect"
	"testing"
)

func TestConvertFlagsToPorts(t *testing.T) {
	tests := map[string]struct {
		ports    []string
		expected []string
	}{
		"Basic": {
			ports:    []string{"1", "100", "300", "100000", "ble"},
			expected: []string{"1", "100", "300"},
		},
		"Range": {
			ports:    []string{"1", "100", "300", "25-28", "100000", "ble"},
			expected: []string{"1", "100", "300", "25", "26", "27", "28"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ConvertFlagToPorts(test.ports)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Validation does not work: expected %v, got %v", test.expected, actual)
			}
		})
	}

}

func TestValidatePort(t *testing.T) {
	tests := map[string]struct {
		port     string
		expected bool
	}{
		"100": {
			port:     "100",
			expected: true,
		},
		"0": {
			port:     "0",
			expected: false,
		},
		"ble": {
			port:     "ble",
			expected: false,
		},
		"100000": {
			port:     "100000",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := validatePort(test.port)

			if actual != test.expected {
				t.Errorf("Validation does not work: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
