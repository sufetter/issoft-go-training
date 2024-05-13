package main

import (
	"log"
	"webcrm/internal/api"
)

func main() {
	err := api.SetMux("8080")
	if err != nil {
		log.Fatal(err)
	}
}
