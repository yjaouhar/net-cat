package util

import (
	"fmt"
	"os"
)

func Port() string {
	port := "8989"
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat [$port]")
		return ""
	}
	if len(args) == 1 {
		port = args[0]
	}
	return port
}
