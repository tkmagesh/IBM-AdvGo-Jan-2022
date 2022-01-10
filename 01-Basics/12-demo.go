package main

import "fmt"

func main() {
	isPrime := memoize(checkPrime)
	fmt.Println(isPrime(96))
	fmt.Println(isPrime(97))
	fmt.Println(isPrime(98))
	fmt.Println(isPrime(99))

	fmt.Println("Processing again....")
	fmt.Println(isPrime(96))
	fmt.Println(isPrime(97))
	fmt.Println(isPrime(98))
	fmt.Println(isPrime(99))
}

func checkPrime(n int) bool {
	fmt.Println("Processing ", n)
	if n <= 1 {
		return false
	}
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func memoize(oper func(n int) bool) func(n int) bool {
	cache := make(map[int]bool)
	return func(n int) bool {
		if _, ok := cache[n]; !ok {
			cache[n] = oper(n)
		}
		return cache[n]
	}
}
