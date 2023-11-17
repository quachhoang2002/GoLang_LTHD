package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}

func main() {
	// Goroutine
	go sayHello("Viet")

	// normal function
	sayHello("Nam")
	time.Sleep(time.Second) // sleep 1s
}
