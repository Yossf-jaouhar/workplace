package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var port = "8989"

var (
	clients     = make(map[net.Conn]string)
	clientsLock sync.Mutex
)

func main() {
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	ls, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("err to lesnning", err)
		return
	}

	for {
		con, err := ls.Accept()
		if err != nil {
			fmt.Println("err to ", err)
			continue
		}

		go handlercon(con)
	}
}

func handlercon(con net.Conn) {

	defer func() {
		clientsLock.Lock()
		delete(clients, con)
		clientsLock.Unlock()
	}()
	
	_, err := con.Write([]byte("your name??"))
	if err != nil {
		return
	}

	buffer := make([]byte, 1024)
	n, err := con.Read(buffer)
	if err != nil {
		return
	}

	username := string(buffer[:n])
	validationMsg := validstring(username)
	for validationMsg != "" {
		if len(clients) > 1 {
			return
		}
		_, err := con.Write([]byte(validationMsg + "\nPlease enter a valid name: "))
		if err != nil {
			return
		}
		n, err = con.Read(buffer)
		if err != nil {
			return
		}
		username = string(buffer[:n])
		validationMsg = validstring(username)
	}

	clientsLock.Lock()
	clients[con] = username
	clientsLock.Unlock()
	fmt.Println("New client:", username)

}

func validstring(name string) string {
	if len(name) > 25 {
		return "Name is too long; maximum 25 characters."
	} else if len(name) < 3 {
		return "Name is too short; minimum 3 characters."
	}
	return ""
}
