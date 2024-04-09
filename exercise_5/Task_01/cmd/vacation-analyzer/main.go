package main

import (
	"bufio"
	"bytes"
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
	s := time.Now()
	dir, err := os.Getwd()
	storagePath := dir + "/storage/"
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	file, err := os.Open(storagePath + "data.csv")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	fileInfo, err := os.Stat(storagePath + "data.csv")
	if err != nil {
		log.Fatal("Error getting file info:", err)
	}

	reader := csv.NewReader(bufio.NewReader(file))

	var (
		vacations        map[string]int
		lineNum          int
		allErrors        []error
		record           []string
		duration         int
		start, end       time.Time
		startErr, endErr error
		errorBuffer      bytes.Buffer
	)

	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("errVal reading file at line: %d %s", lineNum+1, strings.Join(record, ",")))
		}
		lineNum++

		if len(record) != 3 || len(strings.Split(record[0], " ")) != 2 {
			allErrors = append(allErrors, fmt.Errorf("wrong format line: %d %s", lineNum, strings.Join(record, ",")))
		}

		start, startErr = time.Parse(layout, record[1])
		end, endErr = time.Parse(layout, record[2])
		if startErr != nil || endErr != nil {
			allErrors = append(allErrors, fmt.Errorf("wrong date format at line: %d %s", lineNum, strings.Join(record, ",")))
		}

		duration = int(end.Sub(start).Hours() / 24)
		if duration < 0 || duration == 0 {
			allErrors = append(allErrors, fmt.Errorf("wrong date duration errVal at line: %d %s", lineNum, strings.Join(record, ",")))
			continue
		}

		if vacations == nil {
			vacations = make(map[string]int, fileInfo.Size()/int64(len(strings.Join(record, ",")))/2)
		}
		vacations[record[0]] += duration
	}

	if len(allErrors) > 0 {
		for _, errVal := range allErrors {
			_, err := fmt.Fprintln(&errorBuffer, errVal)
			if err != nil {
				log.Fatal("Error writing errors to buffer:", err)
			}
		}

		errorFile, err := os.Create(storagePath + "errors.txt")
		if err != nil {
			log.Fatal("Error creating errVal file:", err)
		}
		defer errorFile.Close()

		_, err = errorFile.Write(errorBuffer.Bytes())
		if err != nil {
			log.Fatal("Error writing errors to file:", err)
		}
	}

	employees := make([]*Employee, 0, len(vacations))
	for k, v := range vacations {
		employees = append(employees, &Employee{k, v})
	}

	sort.Slice(employees, func(i, j int) bool {
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
	defer buf.Flush()
	for _, employee := range employees {
		_, err := fmt.Fprintf(buf, "%s: %d\n", employee.Name, employee.Vacation)
		if err != nil {
			log.Fatal("Error writing to out file:", err)
		}
	}

	fmt.Println(time.Since(s))
}
