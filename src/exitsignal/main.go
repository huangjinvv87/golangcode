package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	go exitSignalDeamon()
	stop()
	select {}
}

func stop() {
	time.Sleep(5 * time.Second)
	fmt.Println("sleep 5 second")
	process, _ := os.FindProcess(os.Getpid())
	process.Signal(syscall.SIGTERM)
}

func clear(desc string) {
	fmt.Println(desc)
}
