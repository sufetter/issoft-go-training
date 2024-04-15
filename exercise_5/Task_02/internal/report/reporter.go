package report

import (
	"log"
	"os"
	"path/filepath"
	"sales-reporter/internal/csvparser"
	"text/template"
	"time"
)

func GenerateReport(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	salesReport, err := csvparser.ParseSalesData(filepath.Join(dir, path))
	if err != nil {
		return err
	}

	tmpl := template.New("report.tmpl").Funcs(template.FuncMap{
		"Add": func(a, b int) int {
			return a + b
		},
	})

	tmpl, err = tmpl.ParseFiles(filepath.Join(dir, "static", "templates", "report.tmpl"))
	if err != nil {
		log.Print("Template parsing error: ", err)
		return err
	}

	reportFile, err := os.Create(filepath.Join(dir, "storage", "report.md"))
	if err != nil {
		log.Print("File creation error: ", err)
		return err
	}
	defer reportFile.Close()

	reportData := struct {
		Date                  string
		TotalRevenue          float64
		HighestRevenueProduct string
		HighestRevenue        float64
		Products              map[string]float64
	}{
		Date:                  time.Now().Format("02 January, 2006"),
		TotalRevenue:          salesReport.TotalRevenue,
		HighestRevenueProduct: salesReport.HighestRevenueProduct,
		HighestRevenue:        salesReport.HighestRevenue,
		Products:              salesReport.ProductSales,
	}

	err = tmpl.Execute(reportFile, reportData)
	if err != nil {
		log.Print("Template execution error: ", err)
		return err
	}

	return nil
}
