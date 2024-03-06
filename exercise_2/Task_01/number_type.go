package main

import "fmt"

func checkNumberType(n int) (string, bool) {
	if n <= 1 {
		return "", false
	}

	count := 0

	for n%2 == 0 {
		n /= 2
		count++
	}

	i := 3
	for i*i <= n {
		if n%i == 0 {
			n /= i
			count++
		} else {
			i += 2
		}
	}

	if n > 2 {
		count++
	}

	if count%2 == 0 {
		return "even", true
	}
	return "odd", true
}

func main() {
	testCases := [...]int{
		5,
		-3,
		2,
		10,
		-1,
	}

	for _, tc := range &testCases {
		numType, flag := checkNumberType(tc)
		fmt.Printf("Result: %v, %v\n", numType, flag)
	}
}
