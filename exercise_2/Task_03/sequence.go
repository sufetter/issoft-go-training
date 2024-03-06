package main

import (
	"fmt"
)

func sequence(nums ...int) []int {
	switch len(nums) {
	case 0:
		return []int{}
	case 1:
		return generateSequence(0, nums[0])
	case 2:
		return generateSequence(nums[0], nums[1])
	default:
		return nums
	}
}

func generateSequence(start, end int) []int {
	if start > end {
		start, end = end, start
	}
	seq := make([]int, 0, end-start+1)
	for ; start <= end; start++ {
		seq = append(seq, start)
	}
	return seq
}

func main() {
	testCases := [...][]int{
		{5},
		{-3},
		{2, 6},
		{10, 20, 30, 40},
		{},
	}
	for _, tc := range testCases {
		fmt.Printf("Result: %v\n", sequence(tc...))
	}
}
