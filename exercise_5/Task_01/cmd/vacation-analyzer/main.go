package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Employee struct {
	Name     string
	Vacation int
}

const layout = "1/2/2006"

func main() {
	dir, err := os.Getwd()
	storagePath := dir + "/storage/"
	if err != nil {
		log.Fatal("Error getting current directory:", err)

	}

	file, err := os.Open(storagePath + "data.csv")
	if err != nil {
		log.Fatal("Error opening file:", err)

	}

	fileInfo, err := os.Stat(storagePath + "data.csv")
	if err != nil {
		log.Fatal("Error getting file info:", err)
	}

	defer file.Close()

	errorFile, err := os.Create(storagePath + "errors.txt")
	if err != nil {
		log.Fatal("Error creating error file:", err)
	}
	defer errorFile.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	var vacations map[string]int
	var lineNum int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(errorFile, "%d %s\n", lineNum+1, strings.Join(record, ","))
		}
		lineNum++

		if len(record) != 3 || len(strings.Split(record[0], " ")) != 2 {
			fmt.Fprintf(errorFile, "%d %s\n", lineNum, strings.Join(record, ","))
		}

		start, err1 := time.Parse(layout, record[1])
		end, err2 := time.Parse(layout, record[2])
		if err1 != nil || err2 != nil {
			fmt.Fprintf(errorFile, "%d %s\n", lineNum, strings.Join(record, ","))
		}

		duration := int(end.Sub(start).Hours() / 24)
		if duration < 0 {
			fmt.Fprintf(errorFile, "%d %s\n", lineNum, strings.Join(record, ","))
		}
		if vacations == nil {
			fmt.Println(len(strings.Join(record, ",")))
			fmt.Println(fileInfo.Size() / int64(len(strings.Join(record, ","))))
			vacations = make(map[string]int, fileInfo.Size()/int64(len(strings.Join(record, ","))))
		}
		vacations[record[0]] += duration
	}

	employees := make([]Employee, 0, len(vacations))
	for k, v := range vacations {
		employees = append(employees, Employee{k, v})
	}

	sort.SliceStable(employees, func(i, j int) bool {
		if employees[i].Vacation == employees[j].Vacation {
			return employees[i].Name < employees[j].Name
		}
		return employees[i].Vacation > employees[j].Vacation
	})

	outFile, err := os.Create(storagePath + "out.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	buf := bufio.NewWriter(outFile)
	for _, employee := range employees {
		fmt.Fprintf(buf, "%s: %d\n", employee.Name, employee.Vacation)
	}
	buf.Flush()
}
