package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input int positive number: ")
	fmt.Scan(&num)

	ternary := 0
	factor := 1
	numCopy := num

	for numCopy > 0 {
		ternary += (numCopy % 3) * factor
		numCopy /= 3
		factor *= 10
	}

	reverse := 0
	numCopy = ternary
	for numCopy > 0 {
		reverse = reverse*10 + numCopy%10
		numCopy /= 10
	}

	if ternary == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Fake")
	}
}
