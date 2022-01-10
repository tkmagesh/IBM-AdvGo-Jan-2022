package main

import "fmt"

func main() {
	loggedAdd := getLoggerOperation(add)
	loggedSubtract := getLoggerOperation(subtract)
	loggedAdd(100, 200)
	loggedSubtract(100, 200)
}

func getLoggerOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		oper(x, y)
		fmt.Println("After invocation")
	}
}

func add(x, y int) {
	result := x + y
	fmt.Println("result = ", result)
}

func subtract(x, y int) {

	result := x - y
	fmt.Println("result = ", result)

}
