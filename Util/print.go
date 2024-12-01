package util

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Prnt(message string, smia string, b bool) {
	defer Mutx.Unlock()
	time := time.Now().Format(time.DateTime)
	Mutx.Lock()
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	m := io.MultiWriter(os.Stdout, f)

	if len(Mp) == 1 && !b {
		m.Write([]byte("[" + time + "]" + "[" + smia + "]:" + message))
	}

	for name, con := range Mp {
		if name != smia {
			if b {
				con.Write([]byte(message))
			} else {
				fmt.Fprintln(con)
				m.Write([]byte("[" + time + "]" + "[" + smia + "]:" + message))
				con.Write([]byte("[" + time + "]" + "[" + smia + "]:" + message))
			}

			fmt.Fprint(con, "["+time+"]"+"["+name+"]:")
		}
	}
}
