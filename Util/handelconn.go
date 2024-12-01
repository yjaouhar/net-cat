package util

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func HandleConn(conn net.Conn) {
	name := ""
	var err error
	p := 0

	for {
		if p == 0 {
			p++
			Mutx.Lock()
		}
		fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
		name, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			Mutx.Unlock()
			return
		}
		if name == "\n" || !Valid(name[:len(name)-1]) {
			continue
		}

		if _, ok := Mp[strings.TrimSpace(name)]; !ok {

			Mp[strings.TrimSpace(name)] = conn
			Mutx.Unlock()
			notify(strings.TrimSpace(name), "joined")
			break
		} else {
			fmt.Fprintf(conn, "[ERROR]:   name already taken\n")
		}
	}

	data, _ := os.ReadFile("logs.txt")
	i := 0
	for {
		time := time.Now().Format(time.DateTime)
		if i == 0 {
			if len(data) != 0 {
				fmt.Fprint(conn, string(data))
			}
			fmt.Fprint(conn, "["+time+"]"+"["+strings.TrimSpace(name)+"]:")
			i++

		} else {
			fmt.Fprint(conn, "["+time+"]"+"["+strings.TrimSpace(name)+"]:")
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if !Valid(strings.TrimSpace(message)) || len(message) == 1 {
			continue
		}
		if err != nil {
			notify(strings.TrimSpace(name), "left")
			return
		}
		Prnt(message, strings.TrimSpace(name), false)

	}
}
