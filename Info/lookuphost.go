package main

import (
	"fmt"  // Package for printing output
	"net"  // Package for network-related functions
)

func main() {
	// Define the hostname for which we want to find IP addresses
	host := "www.youtube.com"

	// Use LookupHost to find IP addresses associated with the hostname
	ips, err := net.LookupHost(host)
	// Check if there was an error during the lookup process
	if err != nil {
		// Print the error message and exit if there's an error
		fmt.Println("Error resolving hostname:", err)
		return
	}

	// Print a header message showing the hostname's IP addresses
	fmt.Println("IP addresses for", host)
	// Loop through each IP address in the list and print it
	for _, ip := range ips {
		fmt.Println(ip) // Print each IP address
	}
}
