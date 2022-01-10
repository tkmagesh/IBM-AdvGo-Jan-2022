package main

import "fmt"

func main() {
	increment := getCounter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func getCounter() func() int { //step - 1
	var no = 0                //step - 2
	increment := func() int { //step - 3
		no++ //step - 4
		return no
	}
	return increment //step - 5
}
