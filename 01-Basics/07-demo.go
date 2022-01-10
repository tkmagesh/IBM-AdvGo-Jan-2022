package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(x int) {
			printNo(x)
		}(i)
	}
	wg.Wait()
}
func printNo(no int) {
	fmt.Println("No = ", no)
	wg.Done()
}
