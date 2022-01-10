package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go f1(wg) //=> scheduled for execution
	f2()
	t := time.Now()
	fmt.Println("Almost end of main")
	wg.Wait()
	fmt.Println("t in main", t)
}

func f1(wg *sync.WaitGroup) {
	t := time.Now()
	fmt.Println("f1 is invoked")
	fmt.Println("t in f1", t)
	wg.Done()
}

func f2() {
	fmt.Println("f2 is invoked")
}
