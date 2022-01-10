package main

import "fmt"

type operation func(int, int) int

func main() {
	//Functions can be assigned to variables
	fn := func() {
		fmt.Println("fn is invoked")
	}
	//fn()

	//functions can be passed as arguments
	exec(fn)

	//Functions can be returned as return values
	adder := getAdder()
	fmt.Println(adder(100, 200))

	//Anonymous functions
	func() {
		fmt.Println("Anonymous function invoked")
	}()
}

func exec(f func()) {
	f()
}

/*
func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getSubtractor() func(int, int) int {
	return func(x, y int) int {
		return x - y
	}
}
*/

func getAdder() operation {
	return func(x, y int) int {
		return x + y
	}
}

func getSubtractor() operation {
	return func(x, y int) int {
		return x - y
	}
}
