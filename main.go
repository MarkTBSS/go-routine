package main

import (
	"fmt"
	"time"
)

func main() {
	// Synchronous or Sequential
	fmt.Println("Hello, world : 1")
	fmt.Println("Hello, world : 2")
	fmt.Println("Hello, world : 3")
	fmt.Println("Hello, world : 4")
	fmt.Println("Hello, world : 5")
	fmt.Println("================")

	// Asynchronous
	go fmt.Println("Hello, world : A")
	go fmt.Println("Hello, world : B")
	go fmt.Println("Hello, world : C")
	go fmt.Println("Hello, world : D")
	go fmt.Println("Hello, world : E")
	time.Sleep(1 * time.Second)
}
