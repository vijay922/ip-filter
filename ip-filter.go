package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

// List of IP ranges to filter
var ipRanges = []string{
	"103.21.244.0/22", "103.22.200.0/22", "103.31.4.0/22", "104.16.0.0/12", "104.16.0.0/13",
	"104.24.0.0/14", "104.64.0.0/10", "108.162.192.0/18", "118.214.0.0/16", "131.0.72.0/22",
	"141.101.64.0/18", "162.158.0.0/15", "172.64.0.0/13", "173.222.0.0/15", "173.245.48.0/20",
	"184.24.0.0/13", "184.50.0.0/15", "184.84.0.0/14", "188.114.96.0/20", "190.93.240.0/20",
	"197.234.240.0/22", "198.41.128.0/17", "2.16.0.0/13", "23.0.0.0/12", "23.192.0.0/11",
	"23.32.0.0/11", "23.64.0.0/14", "23.72.0.0/13", "69.192.0.0/16", "72.246.0.0/15",
	"88.221.0.0/16", "92.122.0.0/15", "95.100.0.0/15", "96.16.0.0/15", "96.6.0.0/15",
}

// Checks if an IP is in any of the given CIDR ranges
func isInRange(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	for _, cidr := range ipRanges {
		_, network, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Printf("Error parsing CIDR %s: %v\n", cidr, err)
			continue
		}
		if network.Contains(parsedIP) {
			return true
		}
	}
	return false
}

func main() {
	// Parse command-line arguments
	outputFile := flag.String("output", "", "Output file for IPs not in the ranges (default: stdout)")
	flag.Parse()

	// Determine input source: file or stdin
	var input *os.File
	if len(flag.Args()) > 0 {
		var err error
		input, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Printf("Error opening input file: %v\n", err)
			return
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	// Open output destination
	var writer *bufio.Writer
	if *outputFile != "" {
		outFile, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer outFile.Close()
		writer = bufio.NewWriter(outFile)
		defer writer.Flush()
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	// Process input
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip == "" {
			continue
		}
		if !isInRange(ip) {
			_, err := writer.WriteString(ip + "\n")
			if err != nil {
				fmt.Printf("Error writing to output: %v\n", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	// Flush output
	writer.Flush()
}
