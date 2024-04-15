package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SalesReport struct {
	ProductSales          map[string]float64
	TotalRevenue          float64
	HighestRevenue        float64
	HighestRevenueProduct string
}

func ParseSalesData(filepath string) (*SalesReport, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		productSales          = make(map[string]float64)
		scanner               = bufio.NewScanner(file)
		totalRevenue          float64
		highestRevenue        float64
		highestRevenueProduct string
	)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		product := strings.Trim(fields[0], "\"")

		price, err := strconv.ParseFloat(fields[1], 64)
		if err != nil || price <= 0 {
			fmt.Println("Error parsing price:", err, " - ", line)
			continue
		}
		units, err := strconv.Atoi(fields[2])
		if err != nil || units <= 0 {
			fmt.Println("Error parsing units:", err, " - ", line)
			continue
		}
		revenue := price * float64(units)
		totalRevenue += revenue
		productSales[product] += revenue

		if productSales[product] > highestRevenue {
			highestRevenue = productSales[product]
			highestRevenueProduct = product
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &SalesReport{
		ProductSales:          productSales,
		TotalRevenue:          totalRevenue,
		HighestRevenue:        highestRevenue,
		HighestRevenueProduct: highestRevenueProduct,
	}, nil
}
