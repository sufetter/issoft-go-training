package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	w := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}

func checkNumberType(n int) (string, bool) {
	if n <= 1 {
		return "", false
	}

	primeFactors := make(map[int]bool)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			primeFactors[i] = true
			n /= i
		}
	}

	if isPrime(n) {
		primeFactors[n] = true
	}

	if len(primeFactors)%2 == 0 {
		return "even", true
	}
	return "odd", true
}

func main() {
	//Данный алгоритм менялся много раз и я не уверен,
	//что на данный момент он работает лучшим образом,
	//но текущая сложность O(sqrt(n))
	testCases := [...]int{
		5,
		-3,
		2,
		10,
		-1,
	}

	s := time.Now()
	for _, tc := range testCases {
		numType, flag := checkNumberType(tc)
		fmt.Printf("Result: %v, %v\n", numType, flag)
	}
	fmt.Printf("%v", time.Since(s))
}
