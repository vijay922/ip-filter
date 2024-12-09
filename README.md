# IP-Filter
IP Filter Tool  A Go script that filters out IPs from a provided list if they belong to specific Cloudflare or Akamai IP ranges. The tool supports input from files or standard input (via pipes) and outputs the filtered list to standard output or a specified file.

## Features

- **Filter Cloudflare and Akamai IPs**: Removes IPs matching predefined ranges.
- **Flexible Input**: Accepts IPs from a file or standard input.
- **Custom Output**: Prints results to the terminal or saves them to a file.
- **Efficient Filtering**: Leverages Go's robust networking capabilities.

## Prerequisites

- [Go](https://go.dev/) (1.17 or later)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/vijay922/ip-filter.git
   cd ip-filter
   go build -o ip-filter
   ```
   # or
2. Install the ip-filter using the Go tool:
```
go install github.com/vijay922/ip-filter@latest
```

# Usage
Run with a file as input
Filter IPs from ips.txt and output the result to the terminal:

```
go run ip-filter.go ips.txt
```

Run with a file and save to an output file
Filter IPs from ips.txt and save the filtered IPs to filtered_ips.txt:
```
go run ip-filter.go ips.txt --output filtered_ips.txt
```
Pipe input via standard input
Use a pipeline to filter IPs:
```
cat ips.txt | ip-filter
```
Pipe input and save to a file
Pipe IPs and save the filtered results to filtered_ips.txt:
```
cat ips.txt | ip-filter --output filtered_ips.txt
```
# Example Input
ips.txt:
```
103.21.244.5
192.168.1.1
104.16.20.5
8.8.8.8
```
# Example Output
Filtered output (if not saved to a file):
```
192.168.1.1
8.8.8.8
```
