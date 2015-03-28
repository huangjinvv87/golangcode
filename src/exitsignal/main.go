package main

import (
	"fmt"
)

var exitChan = make(chan bool, 1)

func main() {
	go exitSignalDeamon()

	select {
	case <-exitChan:
		return
	}
}

func clear(desc string) {
	fmt.Println(desc)
}
