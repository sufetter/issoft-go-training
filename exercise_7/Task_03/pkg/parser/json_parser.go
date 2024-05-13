package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	Dir string
)

func init() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory: ", err)
	}
	Dir = d
}

func Config(d string) error {
	if d != "" {
		Dir = filepath.Join(Dir, d)
		return nil
	}
	return fmt.Errorf("directory value is empty")
}

func WriteJSON(data any, path string) error {
	file, err := os.Create(filepath.Join(Dir, path))
	if err != nil {
		log.Print("Error creating file: ", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		log.Print("Error encoding data: ", err)
		return err
	}

	return nil
}

func ReadJSON(target any, path string) error {
	file, err := os.Open(filepath.Join(Dir, path))
	if err != nil {
		log.Print("Error opening file: ", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		log.Print("Error decoding data: ", err)
		return err
	}

	return nil
}
