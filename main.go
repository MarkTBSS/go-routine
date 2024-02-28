package main

import (
	"fmt"
	"time"
)

func main() {
	startSyncTime := time.Now()
	// Synchronous or Sequential
	func() {
		start1 := time.Now()
		fmt.Println("Start : 1")
		time.Sleep(1 * time.Second)
		fmt.Println("End : 1 = ", time.Since(start1))
	}()
	func() {
		start2 := time.Now()
		fmt.Println("Start : 2")
		time.Sleep(2 * time.Second)
		fmt.Println("End : 2 = ", time.Since(start2))
	}()
	func() {
		start3 := time.Now()
		fmt.Println("Start : 3")
		time.Sleep(4 * time.Second)
		fmt.Println("End : 3 = ", time.Since(start3))
	}()
	fmt.Println("Total Run Time : ", time.Since(startSyncTime))
	fmt.Println("=================")
	time.Sleep(2 * time.Second)
	startAsyncTime := time.Now()
	// Asynchronous
	go func() {
		startA := time.Now()
		fmt.Println("Start : A")
		time.Sleep(1 * time.Second)
		fmt.Println("End : A = ", time.Since(startA))
	}()
	go func() {
		startB := time.Now()
		fmt.Println("Start : B")
		time.Sleep(2 * time.Second)
		fmt.Println("End : B = ", time.Since(startB))
	}()
	func() {
		startC := time.Now()
		fmt.Println("Start : C")
		time.Sleep(4 * time.Second)
		fmt.Println("End : C = ", time.Since(startC))
	}()
	fmt.Println("Total Run Time : ", time.Since(startAsyncTime))
}
