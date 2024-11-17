package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type info struct {
	clients map[net.Conn]string
	mut     sync.Mutex
	FilePath string
}

func main() {
	port := "8989"
	if len(os.Args) == 2 && isValid(os.Args[1]) {
		port = os.Args[1]
	}
	ls, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server listening on port", port)

	serverinfo := info{clients: make(map[net.Conn]string)}

	for {
		con, err := ls.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go serverinfo.HandlerCon(con)
	}
}

func isValid(arg string) bool {
	for _, char := range arg {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
