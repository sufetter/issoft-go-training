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
	case percentageOfSuccess >= 0 && percentageOfSuccess <= 64:
		fmt.Println("F")
	case percentageOfSuccess >= 65 && percentageOfSuccess <= 69:
		fmt.Println("D")
	case percentageOfSuccess >= 70 && percentageOfSuccess <= 79:
		fmt.Println("C")
	case percentageOfSuccess >= 80 && percentageOfSuccess <= 89:
		fmt.Println("B")
	case percentageOfSuccess >= 90 && percentageOfSuccess <= 100:
		fmt.Println("A")
	}
}
