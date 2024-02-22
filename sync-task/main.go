package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	task1()
	task2()
	task3()
	task4()

	fmt.Println("elapsed : ",time.Since(now))
}

func task1()  {
	time.Sleep(1000*time.Millisecond)
	println("task 1 complete")
}

func task2()  {
	time.Sleep(500*time.Millisecond)
	println("task 2 complete")
}

func task3()  {
	println("task 3 complete")
}

func task4()  {
	time.Sleep(100*time.Millisecond)
	println("task 2 complete")
}