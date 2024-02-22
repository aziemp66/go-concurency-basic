package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan bool)
	go func() {
		work()
		done <- true
	}()
	<-done
	fmt.Println("Done Waiting, main Exist")
	fmt.Println("elapsed : ", time.Since(now))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")
}
