# PortScanner

Concurrently scan IP addresses for open ports

## Setup

To use the PortScanner, you must have `go 1.21` installed on your device.
Then, build the executable in the project directory:

```bash
go build .
```

## Usage

`PortScanner [OPTIONS] IP_ADDRESS1 [IP_ADDRESS2 ...]`

OR

`go run scanner.go [OPTIONS] IP_ADDRESS1 [IP_ADDRESS2 ...]`

## Options

- `-p PORTS`
  Specify the ports to scan for each IP address. Ports can be defined as a single port (80), a comma-separated list of ports (80,443), or a range of ports (80-1000). Multiple ranges and individual ports can be combined by separating them with commas (1,80-100,2000). If no ports are specified, PortScanner will scan all ports from 1 - 65.536

## Arguments

- `IP_ADDRESS`
  The IP address(es) to scan. PortScanner accepts an unlimited number of IP addresses as arguments. Each IP address should be separated by space. IP Addresses can be in the form of actual addresses ([0-255].[0-255].[0-255].[0-255]) or as a domain name (ru.is).

## Examples

1. Scan a single IP address on the default port range:

```
PortScanner 192.168.1.1
```

2. Scan multiple IP addresses on specific ports:

```
PortScanner 192.168.1.1 10.0.0.1 -p 22,80,443
```

3. Scan a single IP address on a range of ports:

```
PortScanner 192.168.1.1 -p 100-200
```

4. Scan multiple IP addresses on a combination of specific ports and port ranges:

```
PortScanner 192.168.1.1 10.0.0.1 172.16.0.1 -p 22,80-88,443
```

5. Scan IP addresses from domains:

```
PortScanner ru.is hi.is -p 22,80-88,443
```

## Testing

Utility functions within this program are thoroughly tested to ensure robustness and reliability of the program. To run the tests:

```bash
go test ./...
```
