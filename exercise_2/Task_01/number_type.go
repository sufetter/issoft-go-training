package main

import (
	"fmt"
)

func checkNumberType(n int) (string, bool) {
	if n <= 1 {
		return "", false
	}

	primeFactors := make([]int, 0)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
	}

	if n > 1 {
		primeFactors = append(primeFactors, n)
	}

	fmt.Print(primeFactors)
	if len(primeFactors)%2 == 0 {
		return "even", true
	}
	return "odd", true
}

func main() {
	testCases := [...]int{
		9,
		3,
		15,
		30,
		63018038201,
	}

	for _, tc := range testCases {
		numType, flag := checkNumberType(tc)
		fmt.Printf("Result: %v, %v\n", numType, flag)
	}
}
