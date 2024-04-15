package csvparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Employee struct {
	Name     string
	Vacation int
}

type FileWriter interface {
	WriteToFile(data any, filePath string) error
}

var nameRegex = regexp.MustCompile("^[a-zA-Zа-яА-Я-\\s]+$")

func ParseEmployees(path string, fileName string, layout string, writer FileWriter) error {
	file, err := os.Open(filepath.Join(path, fileName))
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	location, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("error getting location: %v", err)
	}

	var (
		lineNum   int
		allErrors []string
		vacations = make(map[string]int)
		reader    = csv.NewReader(file)
	)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) != 3 || !nameRegex.MatchString(record[0]) {
			allErrors = append(allErrors, fmt.Sprintf("format: %d %s", lineNum+1, strings.Join(record, ",")))
			continue
		}
		lineNum++

		start, startErr := time.ParseInLocation(layout, record[1], location)
		end, endErr := time.ParseInLocation(layout, record[2], location)
		if startErr != nil || endErr != nil {
			allErrors = append(allErrors, fmt.Sprintf("date: %d %s", lineNum, strings.Join(record, ",")))
			continue
		}

		duration := int(end.Sub(start).Hours() / 24)
		if duration <= 0 {
			allErrors = append(allErrors, fmt.Sprintf("duration: %d %s", lineNum, strings.Join(record, ",")))
			continue
		}

		vacations[record[0]] += duration
	}

	err = writer.WriteToFile(allErrors, filepath.Join(path, "errors.txt"))
	if err != nil {
		return fmt.Errorf("error writing errors to file: %v", err)
	}

	employees := make([]*Employee, 0, len(vacations))
	for name, vacationDays := range vacations {
		employees = append(employees, &Employee{Name: name, Vacation: vacationDays})
	}

	sortEmployeesByVacation(employees)

	err = writer.WriteToFile(employees, filepath.Join(path, "out.txt"))
	if err != nil {
		return fmt.Errorf("error writing employees to file: %v", err)
	}

	return nil
}

func sortEmployeesByVacation(employees []*Employee) {
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Vacation == employees[j].Vacation {
			return employees[i].Name < employees[j].Name
		}
		return employees[i].Vacation > employees[j].Vacation
	})
}
