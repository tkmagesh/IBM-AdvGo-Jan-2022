package main

import "fmt"

var z = 10

func main() {
	addClient()
}

func addClient() {
	//result := add(100, 200)
	//replacing the invocation of the function with the result of the function
	result := 300
	fmt.Println("Result = ", result)

	subResult := subtract(100, 200)
	fmt.Println("subtract result = ", subResult)
	z = 1000
	subResult = subtract(100, 200)
	fmt.Println("subtract result = ", subResult)
}

func add(x, y int) int {
	fmt.Println("Processing", x, "and", y) //=> side effect
	return x + y
}

func subtract(x, y int) int {
	return x - y - z
}
