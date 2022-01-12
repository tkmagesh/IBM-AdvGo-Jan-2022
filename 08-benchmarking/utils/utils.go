package utils

import "math"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime2(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i < (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime3(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	no := float64(n)
	upperRange := int(math.Sqrt(no))
	for i := 2; i < upperRange; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
