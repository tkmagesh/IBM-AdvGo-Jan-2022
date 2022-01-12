package utils

import "testing"

type PrimeTestCase struct {
	number         int
	expectedResult bool
	name           string
	actualResult   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []PrimeTestCase{
		{number: 47, expectedResult: true, name: "47 is prime"},
		{number: 44, expectedResult: false, name: "44 is not prime"},
		{number: 23, expectedResult: true, name: "23 is prime"},
		{number: 53, expectedResult: true, name: "53 is prime"},
		{number: 1, expectedResult: false, name: "1 is not prime"},
		{number: 3, expectedResult: true, name: "3 is prime"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actualResult = IsPrime(testCase.number)
			if testCase.actualResult != testCase.expectedResult {
				t.Errorf("Expected %v but got %v", testCase.expectedResult, testCase.actualResult)
			}
		})
	}

}

func Test_IsPrime2(t *testing.T) {
	testCases := []PrimeTestCase{
		{number: 47, expectedResult: true, name: "47 is prime"},
		{number: 44, expectedResult: false, name: "44 is not prime"},
		{number: 23, expectedResult: true, name: "23 is prime"},
		{number: 53, expectedResult: true, name: "53 is prime"},
		{number: 1, expectedResult: false, name: "1 is not prime"},
		{number: 3, expectedResult: true, name: "3 is prime"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actualResult = IsPrime2(testCase.number)
			if testCase.actualResult != testCase.expectedResult {
				t.Errorf("Expected %v but got %v", testCase.expectedResult, testCase.actualResult)
			}
		})
	}

}

func Test_IsPrime3(t *testing.T) {
	testCases := []PrimeTestCase{
		{number: 47, expectedResult: true, name: "47 is prime"},
		{number: 44, expectedResult: false, name: "44 is not prime"},
		{number: 23, expectedResult: true, name: "23 is prime"},
		{number: 53, expectedResult: true, name: "53 is prime"},
		{number: 1, expectedResult: false, name: "1 is not prime"},
		{number: 3, expectedResult: true, name: "3 is prime"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actualResult = IsPrime3(testCase.number)
			if testCase.actualResult != testCase.expectedResult {
				t.Errorf("Expected %v but got %v", testCase.expectedResult, testCase.actualResult)
			}
		})
	}

}
func Benchmark_IsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(47)
	}
}

func Benchmark_IsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime2(47)
	}
}

func Benchmark_IsPrime3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime3(47)
	}
}
