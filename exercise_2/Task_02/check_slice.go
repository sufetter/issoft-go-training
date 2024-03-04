package main

import "fmt"

func where(slice []int, predicate func(int) bool) []int {
	result := make([]int, 0, len(slice))
	for _, num := range slice {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func foreach(slice []int, action func(int)) {
	for _, num := range slice {
		action(num)
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	selected := where(slice, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println("Selected:", selected)

	foreach(slice, func(num int) {
		fmt.Println("Current number:", num)
	})
}
