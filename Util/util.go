package util

import (
	"net"
	"sync"
)

var (
	Mp   map[string]net.Conn
	Mutx sync.Mutex
)
