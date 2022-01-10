package main

import "fmt"

func main() {
	/*
		add(100,200)
		subtract(100,200)
		loggedAdd(100, 200)
		loggedSubtract(100, 200)
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
}

func logOperation(oper func(int, int), x int, y int) {
	fmt.Println("Before invocation")
	oper(x, y)
	fmt.Println("After invocation")
}

func add(x, y int) {
	result := x + y
	fmt.Println("result = ", result)
}

func subtract(x, y int) {

	result := x - y
	fmt.Println("result = ", result)

}
