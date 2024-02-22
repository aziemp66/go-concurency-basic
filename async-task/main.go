package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	done := make(chan bool)
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	// var wg sync.WaitGroup
	// wg.Add(4)
	// go func() {
	// 	defer wg.Done()
	// 	task1()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	task2()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	task3()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	task4()
	// }()

	<-done
	<-done
	<-done
	<-done

	// wg.Wait()

	fmt.Println("elapsed : ", time.Since(now))
}

func task1(done chan bool) {
	time.Sleep(300 * time.Millisecond)
	println("task 1 complete")
	done <- true
}

func task2(done chan bool) {
	time.Sleep(500 * time.Millisecond)
	println("task 2 complete")
	done <- true
}

func task3(done chan bool) {
	println("task 3 complete")
	done <- true
}

func task4(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	println("task 4 complete")
	done <- true
}

// func task1() {
// 	time.Sleep(1000 * time.Millisecond)
// 	println("task 1 complete")
// }

// func task2() {
// 	time.Sleep(500 * time.Millisecond)
// 	println("task 2 complete")
// }

// func task3() {
// 	println("task 3 complete")
// }

// func task4() {
// 	time.Sleep(100 * time.Millisecond)
// 	println("task 4 complete")
// }
