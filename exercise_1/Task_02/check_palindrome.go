package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input int positive number: ")
	fmt.Scan(&num)

	ternary := ""
	for num > 0 {
		ternary = fmt.Sprintf("%d", num%3) + ternary
		num /= 3
	}

	isPalindrome := true
	numTernaryLength := len(ternary)
	for i := 0; i < numTernaryLength/2; i++ {
		if ternary[i] != ternary[numTernaryLength-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Fake")
	}
}
