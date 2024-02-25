package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Input number: ")
	fmt.Scan(&num)

	if num >= 1_000_000_000 {
		fmt.Println("Input error")
		return
	}

	maxDigit := 0
	for num != 0 {
		digit := num % 10
		if digit > maxDigit {
			maxDigit = digit
			if maxDigit == 9 {
				break
			}
		}
		num = num / 10
	}

	fmt.Printf("The biggest one: %d\n", maxDigit)
}
