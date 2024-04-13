package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	path string
}

var config Config

func mustLoad(path string) error {
	if path == "" {
		return fmt.Errorf("config path is empty")
	}
	config = Config{path}
	return nil
}
func WriteJSON(data any) error {
	file, err := os.Create(config.path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(target any) error {
	file, err := os.Open(config.path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&target)
	if err != nil {
		return err
	}

	return nil
}
