package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go add(100, 200, ch)
	f2()
	result := <-ch
	fmt.Println("result = ", result)
	fmt.Println("exiting from main")
}

func add(x, y int, result chan int) {
	time.Sleep(3 * time.Second)
	result <- x + y
}

func f2() {
	fmt.Println("f2 is invoked")
}
