package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv" // Importing strconv package for string conversion
)

const maxEntriesDefault = 100
const apiURL = "https://api.scanthe.net/"

// Data structure to hold packet information
type Packet struct {
    ID         string `json:"id"`
    Timestamp  string `json:"timestamp"`
    SourceIP   string `json:"source_ip"`
    SourcePort string `json:"source_port"`
    DestPort   string `json:"dest_port"`
    Data       string `json:"data"`
}

// Response structure to match the expected JSON response
type Response struct {
    Data []Packet `json:"data"`
}

func printLogo() {
    fmt.Println(`
  _______                    _______ __           ____ __         __
 |     __|.----.---.-.----- |_     _|  |--.-----.|    |  |.-----.|  |_
 |__     ||  __|  _  |     |  |   | |     |  -__||       ||  -__||   _|
 |_______||____|___._|__|__|  |___| |__|__|_____||__|____||_____||____|

`)
}

func fetchData(url string) (*Response, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("request error: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
    }

    var response Response
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return nil, fmt.Errorf("JSON parse error: %w", err)
    }

    return &response, nil
}

func main() {
    printLogo()

    // Determine how many entries to display (default to max 100)
    maxEntries := maxEntriesDefault
    if len(os.Args) > 1 {
        var err error
        maxEntries, err = strconv.Atoi(os.Args[1]) // Convert command-line argument to integer
        if err != nil || maxEntries < 1 || maxEntries > 100 {
            fmt.Println("Please enter a number between 1 and 100.")
            os.Exit(1)
        }
    }

    // Fetch the data
    response, err := fetchData(apiURL)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Display the relevant parts of the JSON
    for i, packet := range response.Data {
        if i >= maxEntries {
            break // Stop if we reach the maximum entries
        }
        fmt.Printf("ID: %s\n", packet.ID)
        fmt.Printf("Timestamp: %s\n", packet.Timestamp)
        fmt.Printf("Source IP: %s\n", packet.SourceIP)
        fmt.Printf("Source Port: %s\n", packet.SourcePort)
        fmt.Printf("Destination Port: %s\n", packet.DestPort)
        fmt.Printf("Data: %s\n", packet.Data)
        fmt.Println("----------")
    }
}
