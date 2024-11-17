package functions

import (
	"fmt"
	"net"
)

func (a *info) HandlerCon(con net.Conn) {
	defer func() {
		a.mut.Lock()
		delete(a.clients, con)
		a.mut.Unlock()
		con.Close()
	}()

	buffer := make([]byte, 1024)

	_, err := con.Write([]byte("enter your name: "))
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := con.Read(buffer)
	if err != nil {
		return
	}

	name := string(buffer[:n])
	help := false
	ret := isValidName(name)
	if ret == "" {
		help = true
		name = name[:len(name)-1]
		a.mut.Lock()
		a.clients[con] = name
		a.mut.Unlock()
	}
	for !help {

		_, err := con.Write([]byte(ret))
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := con.Read(buffer)
		if err != nil {
			return
		}

		name := string(buffer[:n])
		ret := isValidName(name)
		if ret == "" {
			help = true
			name = name[:len(name)-1]
			a.mut.Lock()
			a.clients[con] = name
			a.mut.Unlock()

		}

	}

	Chat(con)
}

func isValidName(name string) string {
	if len(name) < 3 {
		return "false"
	} else if len(name) > 25 {
		return "false"
	}

	return ""
}
