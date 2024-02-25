package main

import "fmt"

func main() {
	var lowBorder, highBorder int

	fmt.Print("Enter two integers (separated by a space): ")
	fmt.Scanf("%d %d ", &lowBorder, &highBorder)

	if lowBorder > highBorder {
		lowBorder, highBorder = highBorder, lowBorder
	}

	for i := lowBorder; i <= highBorder; i++ {
		count := 0
		num := i
		for num != 0 {
			count += num & 1 // 0 || 1
			num >>= 1        // Переходим к следующему биту (num / 2)
		}
		if count == 4 {
			fmt.Printf("%d - %b\n", i, i)
		}
	}
}
