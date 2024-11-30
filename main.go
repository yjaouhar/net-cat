package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

var (
	sl   []net.Conn
	mp   map[net.Conn]string
	mutx sync.Mutex
)

func main() {
	mp = make(map[net.Conn]string)
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go Hand(con)

	}

}

func Hand(conn net.Conn) {

	fmt.Fprint(conn, "\n [entre name ]:")
	name, _ := bufio.NewReader(conn).ReadString('\n')
	mutx.Lock()
	mp[conn] = name[:len(name)-1]
	sl = append(sl, conn)
	mutx.Unlock()
	data, _:=os.ReadFile("logs.txt")
	i:=0
	for {
		if i==0 {
			time := time.Now().Format(time.DateTime)
		fmt.Fprint(conn, "["+time+"]"+"["+name[:len(name)-1]+"]:")
		fmt.Fprint(conn, string(data))
		i++
		
		}else{
		time := time.Now().Format(time.DateTime)
		fmt.Fprint(conn, "["+time+"]"+"["+name[:len(name)-1]+"]:")
	}
		messge, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			exit(conn)
			return
		}
		Prnt(messge, conn)
		
	}

}

func Prnt(message string, conn net.Conn) {
	defer mutx.Unlock()
	time := time.Now().Format(time.DateTime)
	mutx.Lock()
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	for _, v := range sl {
		if v != conn {
			m := io.MultiWriter(os.Stdout, v,f)
			fmt.Fprintln(v)
			m.Write([]byte("["+time+"]"+"["+mp[conn]+"]:"+message ))
			fmt.Fprint(v, "["+time+"]"+"["+mp[v]+"]:")
		}

	}

}

func exit(conn net.Conn) {
	defer mutx.Unlock()
	mutx.Lock()
	Prnt(fmt.Sprintf("\nrah %s khraj\n", mp[conn]), conn)
	delete(mp, conn)
	fmt.Println(mp)
}
