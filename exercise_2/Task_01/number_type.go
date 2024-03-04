package main

import "fmt"

func checkNumberType(n int) (string, bool) {
	if n <= 1 {
		return "invalid", false
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
	fmt.Println("Input number:")
	var number int
	fmt.Scanf("%d", &number)
	classification, isValid := checkNumberType(number)
	fmt.Printf("Classification: %s, Flag: %v\n", classification, isValid)
}
