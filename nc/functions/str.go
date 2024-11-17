package functions

import (
	"net"
	"sync"
)

type info struct {
	clients  map[net.Conn]string
	mut      sync.Mutex
	FilePath string
}
