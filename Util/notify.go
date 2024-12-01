package util

import (
	"fmt"
)

func notify(name string, s string) {
	Prnt(fmt.Sprintf("\n%v has %v our chat...\n", name, s), name, true)
	if s == "left" {
		defer Mutx.Unlock()
		Mutx.Lock()
		delete(Mp, name)
	}
}
