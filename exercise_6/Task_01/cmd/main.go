package main

import (
	"flag"
	"log"
	"wordfreq/internal/wordcounter"
)

// Предполагается, что файл содержит только слова
// в корректном формате, иначе можно использовать регулярные
// выражения для проверки

func main() {
	// go run ./cmd/main.go -f text.txt -t 200 (по аналогии и другие задания)

	var (
		filePath string
		topWords int
	)
	flag.StringVar(&filePath, "file", "", "path to the text file")
	flag.StringVar(&filePath, "f", "", "file flag shortcut")
	flag.IntVar(&topWords, "top", 10, "number of top words to display")
	flag.IntVar(&topWords, "t", 10, "top flag shortcut")
	flag.Parse()

	if filePath == "" {
		log.Fatal("No file path provided")
	}

	words, err := wordcounter.TopWords(filePath, topWords)
	if err != nil {
		log.Println("Failed to get top words:")
		log.Fatal(err)
	}

	log.Println("Most frequent words:")
	for _, word := range words {
		log.Println(word)
	}
}
