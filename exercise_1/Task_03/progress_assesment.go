package main

import "fmt"

func main() {
	var maxPoints, numTasks int
	fmt.Print("Input maximum number of points: ")
	fmt.Scanln(&maxPoints)

	fmt.Print("Input number of tasks: ")
	fmt.Scanln(&numTasks)

	points := 0
	totalPoints := 0
	for i := 0; i < numTasks; i++ {
		fmt.Printf("Enter the number of points for the task %d: ", i+1)
		fmt.Scanln(&points)
		totalPoints += points
	}

	percentageOfSuccess := float64(totalPoints) / float64(maxPoints*numTasks) * 100
	switch {
	case percentageOfSuccess < 65:
		fmt.Println("F")
	case percentageOfSuccess < 70:
		fmt.Println("D")
	case percentageOfSuccess < 80:
		fmt.Println("C")
	case percentageOfSuccess < 90:
		fmt.Println("B")
	default:
		fmt.Println("A")
	}
}
