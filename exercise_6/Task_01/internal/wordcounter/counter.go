package wordcounter

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

func TopWords(filePath string, topWords int) ([]Word, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)
		for _, word := range words {
			wordMap[word]++
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	words := make([]Word, 0, len(wordMap))
	for word, count := range wordMap {
		words = append(words, Word{word: word, count: count})
	}

	sort.SliceStable(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})

	if topWords < 1 {
		log.Printf("Invalid top words count: %d\nUsing default value", topWords)
		topWords = 10
	}
	if topWords > len(words) {
		log.Printf("Invalid top words count: %d\nUsing len(words) as value", topWords)
		topWords = len(words)
	}

	return words[:topWords], nil
}
