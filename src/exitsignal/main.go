package main

import (
	"fmt"
)

var exitChan = make(chan bool, 1)

func main() {
	go goSignalDeamon()

	select {
	case <-exitChan:
		return
	}
}

func clear(desc string) {
	fmt.Println(desc)
}
