package main

import (
	"fmt"
)

func main() {
	go exitSignalDeamon()
	select {}
}

func clear(desc string) {
	fmt.Println(desc)
}
