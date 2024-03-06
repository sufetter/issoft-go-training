package main

import (
	"fmt"
)

func group(inputMap map[byte]string) map[byte][]string {
	resultMap := make(map[byte][]string, 10)

	for key, value := range inputMap {
		lastDigit := key % 10
		resultMap[lastDigit] = append(resultMap[lastDigit], value)
	}

	return resultMap
}

func main() {
	inputMap := map[byte]string{
		11: "red",
		51: "green",
		22: "blue",
		33: "yellow",
		44: "orange",
		15: "purple",
		92: "black",
	}

	result := group(inputMap)

	fmt.Println("Result:")
	for key, value := range result {
		fmt.Printf("%d: %v\n", key, value)
	}
}
