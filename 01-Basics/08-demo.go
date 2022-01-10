package main

import "fmt"

func main() {
	loggedAdd(100, 200)
	loggedSubtract(100, 200)
}

func loggedAdd(x, y int) {
	fmt.Println("Before invocation")
	add(x, y)
	fmt.Println("After invocation")
}

func loggedSubtract(x, y int) {
	fmt.Println("Before invocation")
	subtract(x, y)
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
