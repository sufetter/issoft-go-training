package wordcounter

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

//Буду честен, потратил много времени
//на отсутствие гонок памяти, не уверен,
//что данный алгоритм оптимален, но он быстрее
//одно поточного по моим замерам примерно в 2-3 раза.
//Проверял гонки через go run --race

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

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	goroutineCount, err := calculateGoroutineCount(fileStat.Size())
	if err != nil {
		return nil, err
	}

	var wordCounts sync.Map

	var wg sync.WaitGroup
	wg.Add(goroutineCount)

	lines := make(chan string, goroutineCount)

	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	for i := 0; i < goroutineCount; i++ {
		go func() {
			defer wg.Done()
			wordMap := make(map[string]int)
			for line := range lines {
				line = strings.ToLower(line)
				words := strings.Fields(line)
				for _, word := range words {
					wordMap[word]++
				}
			}
			for word, count := range wordMap {
				wordCounts.LoadOrStore(word, 0)
				newCount, _ := wordCounts.Load(word)
				wordCounts.Store(word, newCount.(int)+count)
			}
		}()
	}

	wg.Wait()

	var words []Word
	wordCounts.Range(func(key, value interface{}) bool {
		words = append(words, Word{word: key.(string), count: value.(int)})
		return true
	})

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

func calculateGoroutineCount(fileSize int64) (int, error) {
	goroutineCount := 1

	// Assuming files are mostly small
	const smallFileSize = 100 * 1024

	switch {
	case fileSize < smallFileSize:
	case fileSize < smallFileSize*100:
		goroutineCount = 2
	case fileSize < smallFileSize*1000:
		goroutineCount = 4
	default:
		goroutineCount = 8
	}

	return goroutineCount, nil
}
