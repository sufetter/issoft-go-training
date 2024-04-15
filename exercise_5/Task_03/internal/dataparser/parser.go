package dataparser

import (
	"bufio"
	"fmt"
	"knn/internal/entities"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ParseDataFile(fileName string) ([]entities.Object, error) {
	filePath := filepath.Join(".", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var objects []entities.Object
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			log.Printf("Warning: invalid data format at line %d\n", lineNumber)
			continue
		}

		x, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			log.Printf("Warning: %v at line %d\n", err, lineNumber)
			continue
		}

		y, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		if err != nil {
			log.Printf("Warning: %v at line %d\n", err, lineNumber)
			continue
		}

		objects = append(objects, entities.Object{Name: parts[0], X: x, Y: y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(objects) == 0 {
		return nil, fmt.Errorf("no data to classify")
	}

	return objects, nil
}

func ParseCoordinates(text string) (float64, float64, error) {
	parts := strings.Fields(text)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input. Please enter two characteristics separated by space")
	}
	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input for X coordinate: %v", err)
	}
	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input for Y coordinate: %v", err)
	}
	if x < 0 || y < 0 {
		return 0, 0, fmt.Errorf("coordinates must be non-negative")
	}
	return x, y, nil
}
