package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

/*
clean the logs
split code and comments and change var names
*/
var (
	//sl   []net.Conn
	mp   map[string]net.Conn
	mutx sync.Mutex
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

func main() {
	port := Port()
	if port == "" {
		return
	}
	mp = make(map[string]net.Conn)
	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		//os.Stderr.WriteString("failed to run server\n")
		fmt.Fprintf(os.Stderr, "[ERROR]:   %v\n", err)
		return
	}
	fmt.Println("Listening on the port " + port)

	for len(mp) < 10 {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR]:   %v\n", err)
			return
		}
		go HandleConn(conn)

	}

}

func HandleConn(conn net.Conn) {
	defer mutx.Unlock()
	name := ""
	var err error

	for {

		fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
		name, err = bufio.NewReader(conn).ReadString('\n')
		// if !ValidMsg(strings.TrimSpace(name)){
		// 	continue
		// }
		if err != nil {
			return
		}
		if name == "\n" {
			continue
		}
		mutx.Lock()
		if _, ok := mp[strings.TrimSpace(name)]; !ok {
			mp[strings.TrimSpace(name)] = conn
			mutx.Unlock()
			break
		} else {
			fmt.Fprintf(conn, "[ERROR]:   name already taken\n")
		}
		//mp[strings.TrimSpace(name)] = conn
		//sl = append(sl, conn)
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
			//time := time.Now().Format(time.DateTime)
			fmt.Fprint(conn, "["+time+"]"+"["+strings.TrimSpace(name)+"]:")
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			exit(strings.TrimSpace(name))
			return
		}
		fmt.Println([]byte(message), []byte(strings.TrimSpace(message)))
		if ValidMsg(strings.TrimSpace(strings.TrimSpace(message))) {
			Prnt(message, strings.TrimSpace(name))
		}

	}

}

func ValidMsg(s string) bool {
	if s == "" {
		return false
	}
	
	for _, v := range s {
		if !unicode.IsPrint(v) {
			return false
		}
	}
	return true
}

func Prnt(message string, smia string) {
	defer mutx.Unlock()
	time := time.Now().Format(time.DateTime)
	mutx.Lock()
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	for name, con := range mp {
		if len(mp) == 1 {
			m := io.MultiWriter(os.Stdout, f)
			//fmt.Fprintln(con)
			m.Write([]byte("[" + time + "]" + "[" + smia + "]:" + message))
		} else {

			if name != smia {
				m := io.MultiWriter(os.Stdout, mp[name], f)
				fmt.Fprintln(con)
				m.Write([]byte("[" + time + "]" + "[" + smia + "]:" + message))
				fmt.Fprint(con, "["+time+"]"+"["+name+"]:")
			}
		}

	}

}

func exit(name string) {
	defer mutx.Unlock()
	mutx.Lock()
	Prnt(fmt.Sprintf("\nrah %s khraj\n", mp[name]), name)
	delete(mp, name)
	fmt.Println(mp)
}
