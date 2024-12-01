package main

import (
	"fmt"
	"net"
	"os"

	util "net-cat/Util"
)

/*
clean the logs
split code and comments and change var names
*/

func main() {
	port := util.Port()
	if port == "" {
		return
	}
	util.Mp = make(map[string]net.Conn)
	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]:   %v\n", err)
		return
	}
	fmt.Println("Listening on the port " + port)

	for {
		if len(util.Mp) < 10 {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Fprintf(os.Stderr, "[ERROR]:   %v\n", err)
				return
			}
			go util.HandleConn(conn)
		} else {
			continue
		}
	}
}
