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
		11: "BSTU",
		51: "Light saber",
		22: "Jerry",
		33: "Gigachad",
		44: "Sonic",
		15: "Donald Trump",
		92: "Tom",
	}

	result := group(inputMap)

	fmt.Println("Result:")
	for key, value := range result {
		fmt.Printf("%d: %v\n", key, value)
	}
}
