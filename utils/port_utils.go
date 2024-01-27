package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertFlagToPorts(ports []string) []string {
	validPorts := []string{}

	if len(ports) == 1 && ports[0] == "" {
		for i := 1; i <= 65535; i++ {
			validPorts = append(validPorts, strconv.Itoa(i))
		}
		return validPorts
	}

	for _, port := range ports {
		if strings.Contains(port, "-") {
			portRange := convertRangeToPorts(port)
			validPorts = append(validPorts, portRange...)

		}
		if validatePort(port) {
			validPorts = append(validPorts, port)
		}
	}

	return validPorts
}

func validatePort(portStr string) bool {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		errMsg := fmt.Errorf("invalid port: %s is not a number", portStr)
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
		return false
	}
	if port < 1 || port > 65535 {
		errMsg := fmt.Errorf("invalid port: %d is out of range (1-65535)", port)
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
		return false
	}
	return true
}

func convertRangeToPorts(rangeStr string) []string {
	parts := strings.Split(rangeStr, "-")
	result := []string{}
	if len(parts) != 2 {
		err := fmt.Errorf("invalid range format")
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return result
	}

	// Convert strings to integers
	low, errLow := strconv.Atoi(parts[0])
	high, errHigh := strconv.Atoi(parts[1])
	if errLow != nil || errHigh != nil {
		err := fmt.Errorf("range must be numeric")
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return result
	}

	// Check if the range is valid
	if low < 1 || low > 65535 || high < 1 || high > 65535 || low > high {
		err := fmt.Errorf("invalid port range")
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return result
	}

	// Generate the slice of strings
	for i := low; i <= high; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
