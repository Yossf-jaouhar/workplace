package main

import (
    "context"
    "fmt"
    "net"
    "time"
)

func main() {
    // Set the deadline for the context to 5 seconds from the current time
    deadline := time.Now().Add(5 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel() // Call cancel() after the operation to release resources

    // Open a TCP connection to a specific server (in this example, "httpbin.org" on port 80)
    d := net.Dialer{}
    conn, err := d.DialContext(ctx, "tcp", "httpbin.org:80")
    if err != nil {
        fmt.Println("Connection failed:", err)
        return
    }
    defer conn.Close() // Ensure the connection is closed after the operation
    fmt.Println("Connected successfully!")

    // Send an HTTP request manually over the open connection
    request := "GET / HTTP/1.1\r\nHost: httpbin.org\r\n\r\n"
    _, err = conn.Write([]byte(request))
    if err != nil {
        fmt.Println("Failed to send request:", err)
        return
    }

    // Read the response from the server
    buffer := make([]byte, 4096) // Set the temporary memory size for reading
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Failed to read response:", err)
        return
    }

    // Display the received response
    fmt.Println("Response from server:\n", string(buffer[:n]))
}
