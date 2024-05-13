package parser

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func TXTFile[T any](path string, parseFunc func(string) (T, error)) ([]T, error) {
	var err error
	var parsed T

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []T
	var errorLines []string
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		parsed, err = parseFunc(line)
		if err != nil {
			log.Printf("Line %d: %v", lineNumber, err)
			errorLines = append(errorLines, "Line "+strconv.Itoa(lineNumber)+": "+line)
			continue
		}
		result = append(result, parsed)
		lineNumber++
	}

	if err = scanner.Err(); err != nil {
		log.Printf("Error parsing: %v", err)
		return nil, err
	}

	if len(errorLines) > 0 {
		if err = writeErrorsToFile(filepath.Join("storage", "errors.txt"), errorLines); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func writeErrorsToFile(filename string, errorLines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range errorLines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}
