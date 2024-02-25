package main

import "fmt"

func main() {
	var word string
	fmt.Print("Input word: ")
	fmt.Scan(&word)

	isPalindrome := true
	wordLength := len(word)
	for i := 0; i < wordLength/2; i++ {
		if word[i] != word[wordLength-1-i] {
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
