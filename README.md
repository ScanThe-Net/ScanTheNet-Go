# ScanTheNet-Go
Go API Data Fetcher.

This project is a Go application that retrieves and displays data from our API endpoint. 

## Features

- **API Interaction**: Makes GET requests to our API endpoint and retrieves JSON data.
- **Dynamic Entry Limit**: Users can specify the maximum number of data entries to display (up to 100).

## Requirements

- **go**

## Installation

### Install Dependencies

For Debian-based systems, you can install the required dependencies:

```
apt install golang-go

or

wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

Clone the Repository

```
git clone https://github.com/ScanThe-Net/ScanTheNet-Go.git
cd ScanTheNet-Go
```

Usage without build

Run it from the command line. Optionally, you can specify the maximum number of entries:

```
go run main.go [maxEntries]
```

Build

```
go build -o scanthenet main.go
```

Usage with build

Run it from the command line. Optionally, you can specify the maximum number of entries:

```
./scanthenet [maxEntries]
```

Replace [maxEntries] with a number between 1 and 100. If not specified, the default value is 100.

Example Output

When executed successfully, it will display results in the following format:

```
      _______                    _______ __           ____ __         __
     |     __|.----.---.-.----- |_     _|  |--.-----.|    |  |.-----.|  |_
     |__     ||  __|  _  |     |  |   | |     |  -__||       ||  -__||   _|
     |_______||____|___._|__|__|  |___| |__|__|_____||__|____||_____||____|

ID: "5702635"
Timestamp: "2024-11-17 09:26:45"
Source IP: "51.91.110.49"
Source Port: "41756"
Destination Port: "22"
Data: "Flags [S], seq 2981374044, win 29200, options [mss 1400,sackOK,TS val 1400581321 ecr 0,nop,wscale 7], length 0"
----------
```
