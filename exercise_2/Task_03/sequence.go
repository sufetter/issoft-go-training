package main

import "fmt"

func sequence(nums ...int) []int {
	switch len(nums) {
	case 0:
		return []int{}
	case 1:
		n := nums[0]
		if n >= 0 {
			return generateSequence(0, n+1)
		} else {
			return generateSequence(-n, 1)
		}
	case 2:
		a, b := nums[0], nums[1]
		if a <= b {
			return generateSequence(a, b+1)
		} else {
			return generateSequence(b, a+1)
		}
	default:
		return nums
	}
}

func generateSequence(start, end int) []int {
	result := make([]int, end-start)
	for i := 0; i < len(result); i++ {
		result[i] = start + i
	}
	return result
}

func main() {
	fmt.Println(sequence(5))
	fmt.Println(sequence(-3))
	fmt.Println(sequence(2, 6))
	fmt.Println(sequence(10, 20, 30, 40))
	fmt.Println(sequence())
}
