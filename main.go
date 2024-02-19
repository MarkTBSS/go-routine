package main

import (
	"fmt"
	"time"
)

func main() {
	startSyncTime := time.Now()
	// Synchronous or Sequential
	func() {
		fmt.Println("Start : 1")
		time.Sleep(1 * time.Second)
		fmt.Println("End : 1")
	}()
	func() {
		fmt.Println("Start : 2")
		time.Sleep(2 * time.Second)
		fmt.Println("End : 2")
	}()
	func() {
		fmt.Println("Start : 3")
		time.Sleep(4 * time.Second)
		fmt.Println("End : 3")
	}()
	fmt.Println("Total Run Time : ", time.Since(startSyncTime))
	fmt.Println("================")
	time.Sleep(2 * time.Second)
	startAsyncTime := time.Now()
	// Asynchronous
	go func() {
		fmt.Println("Start : A")
		time.Sleep(1 * time.Second)
		fmt.Println("End : A")
	}()
	go func() {
		fmt.Println("Start : B")
		time.Sleep(2 * time.Second)
		fmt.Println("End : B")
	}()
	func() {
		fmt.Println("Start : C")
		time.Sleep(4 * time.Second)
		fmt.Println("End : C")
	}()
	fmt.Println("Total Run Time : ", time.Since(startAsyncTime))
}
