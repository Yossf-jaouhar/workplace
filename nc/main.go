package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var port = "8989"

var (
	clients      = make(map[net.Conn]string)
	mut          sync.Mutex
	chatFilePath = "chat.log"
)

func main() {
	// Use a custom port if provided
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	ls, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// Handle graceful shutdown
	done := make(chan bool)

	go func() {
		ls.Close()

		done <- true
	}()

	fmt.Println("Server is listening on port:", port)

	for {
		con, err := ls.Accept()
		if err != nil {
			if done != nil {
				break
			}
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle connection in a separate goroutine
		go HandlerCon(con)
	}

	<-done
}

// HandlerCon handles a new client connection
func HandlerCon(con net.Conn) {
	defer func() {
		mut.Lock()
		name := clients[con]
		delete(clients, con)
		mut.Unlock()
		con.Close()

		Send("has left the chat.", con, name)
	}()

	buffer := make([]byte, 1024)

	_, err := con.Write([]byte("Enter your name: "))
	if err != nil {
		return
	}

	var clientNm string
	var clientName string
	for {
		n, err := con.Read(buffer)
		if err != nil {
			return
		}

		clientNm = string(buffer[:n])
		clientName = strings.Trim(clientNm, " ")
		msg := ValidateName(clientName)

		if msg == "" {
			break
		} else {
			con.Write([]byte(msg + "\n"))
		}
	}

	fmt.Println(clientName)

	name := clientName[:len(clientName)-1]

	fmt.Println(name)

	mut.Lock()
	clients[con] = name
	mut.Unlock()

	Chat(con, name)
}

// ValidateName validates the client's name
func ValidateName(name string) string {
	if len(name) > 25 {
		return "Name is too long; maximum 25 characters."
	} else if len(name) < 3 {
		return "Name is too short; minimum 3 characters."
	}
	return ""
}

func Chat(con net.Conn, clientName string) {
	buffer := make([]byte, 1024)
	message := ""

	if len(clients) > 0 {
		SendChatHistory(con)
	}

	for {
		n, err := con.Read(buffer)
		if err != nil {
			return
		}

		message = string(buffer[:n])
		AppendToFile(message, clientName)
		Send(message, con, clientName)
	}
}

func Send(ms string, con net.Conn, clientName string) {
	mss := strings.TrimSpace(ms)
	for client := range clients {
		if client != con {
			_, err := client.Write([]byte(clientName + ": " + mss))
			if err != nil {
				return
			}
		}
	}
}

func AppendToFile(mesg string, clientName string) {
	message := strings.Trim(mesg, " ")
	file, err := os.OpenFile(chatFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%s: %s\n", clientName, message))
}

// SendChatHistory reads the chat log file and sends its content to a new client.
func SendChatHistory(con net.Conn) {
	_, err := os.Stat(chatFilePath)
	if os.IsNotExist(err) {
		return
	}

	file, err := os.Open(chatFilePath)
	if err != nil {
		fmt.Println("Error opening chat log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		con.Write([]byte(line + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading chat log file:", err)
	}
}
