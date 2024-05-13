package main

import (
	"belredis/internal/chache"
	"fmt"
	"time"
)

func main() {
	c, err := cache.NewCache(3)
	if err != nil {
		fmt.Println("Error creating cache:", err)
		return
	}

	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}

	for key, value := range data {
		if err := c.Set(key, value); err != nil {
			fmt.Printf("Error setting data for key %s: %v\n", key, err)
		}
	}

	for key := range data {
		value, err := c.Get(key)
		if err != nil {
			fmt.Printf("Error getting data for key %s: %v\n", key, err)
			continue
		}
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	time.Sleep(5 * time.Second)

	err = c.Set("gigakey", "gigachad")
	if err != nil {
		fmt.Printf("Error setting gigachad")
	}
	time.Sleep(499 * time.Millisecond)

	fmt.Println("\nAfter 10 seconds:")
	for key := range data {
		value, err := c.Get(key)
		if err != nil {
			fmt.Printf("Error getting data for key %s: %v\n", key, err)
			continue
		}
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
	item, err := c.Get("gigakey")
	if err != nil {
		fmt.Printf("Error getting gigacad: %v\n", err)
		return
	}
	fmt.Println(item)
}
