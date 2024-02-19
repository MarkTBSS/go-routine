package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Set log flags to include nanosecond timestamp
	startSyncTime := time.Now()
	// Synchronous or Sequential
	log.Println("Asynce : Starting")
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("1")
	}()
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("2")
	}()
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("3")
	}()
	log.Println("Asynce : Finished")
	fmt.Println("Total Run Time : ", time.Since(startSyncTime))
	fmt.Println("================")
	time.Sleep(1 * time.Second)
	startAsyncTime := time.Now()
	// Asynchronous
	log.Println("Synce : Starting")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("A")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("B")
	}()
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("C")
	}()
	log.Println("Synce : Finished")
	fmt.Println("Total Run Time : ", time.Since(startAsyncTime))
}
