package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, ch, wg)
	f2()
	wg.Wait()
	fmt.Println("exiting from main")
}

func add(x, y int, result chan int, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	result <- x + y
	wg.Done()
}

func f2() {
	fmt.Println("f2 is invoked")
}
