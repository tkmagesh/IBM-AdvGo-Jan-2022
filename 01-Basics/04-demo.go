package main

import (
	"fmt"
	"time"
)

func main() {
	ch := add(100, 200)
	f2()
	ch <- 200
	result := <-ch
	fmt.Println("result = ", result)
	fmt.Println("exiting from main")
}

func add(x, y int) chan int {
	result := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		result <- x + y
	}()
	return result
}

func f2() {
	fmt.Println("f2 is invoked")
}
