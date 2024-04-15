package main

import (
	"log"
	"path/filepath"
	"sales-reporter/internal/report"
)

func main() {
	err := report.GenerateReport(filepath.Join("storage", "sales.csv"))
	if err != nil {
		log.Fatal("Error generating report:", err)
	}
}
