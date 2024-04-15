package main

import (
	"bufio"
	"fmt"
	"knn/internal/classifier"
	"knn/internal/dataparser"
	"knn/internal/entities"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getInputCoordinates() (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter two characteristics: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	if strings.TrimSpace(text) == "exit" {
		return -1, -1, fmt.Errorf("exit command received")
	}
	return dataparser.ParseCoordinates(text)
}

func main() {
	//Пока у меня мало знаний относительно
	//архитектуры проектов на Go, поэтому
	//не уверен, что везде все правильно разместил

	objects, err := dataparser.ParseDataFile(filepath.Join("storage", "data.txt"))
	if err != nil {
		log.Fatal("Error parse file", err)
	}

	for {
		x, y, err := getInputCoordinates()
		if x == -1 && y == -1 {
			fmt.Println(err)
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("The object is:", classifier.Classify(objects, entities.Object{X: x, Y: y}, 3))
	}
}
