package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vacation-analyzer/internal/csvparser"
	"vacation-analyzer/internal/utils/writer"
)

const layout = "1/2/2006"

func GetStoragePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %v", err)
	}
	return filepath.Join(dir, "storage"), nil
}

func main() {
	//Пожалуйста, уточняйте, что вы имеете в
	//виду, когда пишете - "некорректные данные"
	//Слегка трудно понять, что именно это
	//значит ситуации с .csv файлами.

	//Я совсем немного изменил ваш .csv файл
	//добавив туда пару ошибок которых там
	//изначально не было для тестирования всех
	//сценариев.

	storagePath, err := GetStoragePath()
	if err != nil {
		log.Fatalf("Error getting storage path: %v", err)
	}

	w := &writer.FileWriter{}
	err = csvparser.ParseEmployees(storagePath, "data.csv", layout, w)
	if err != nil {
		log.Fatalf("Error handling CSV: %v", err)
	}
}
