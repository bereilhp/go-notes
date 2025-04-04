package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := range 5 {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello from Goroutine!") // Run in a separate thread
	printMessage("Hello from Main!")         // Run in main thread
}
