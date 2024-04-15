package writer

import (
	"bufio"
	"fmt"
	"os"
	"vacation-analyzer/internal/csvparser"
)

type FileWriter struct{}

func (w *FileWriter) WriteToFile(data any, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	defer buf.Flush()

	switch v := data.(type) {
	case []string:
		for _, line := range v {
			_, err := fmt.Fprintln(buf, line)
			if err != nil {
				return fmt.Errorf("error writing data to file: %v", err)
			}
		}
	case []*csvparser.Employee:
		for _, emp := range v {
			_, err := fmt.Fprintf(buf, "%s: %d\n", emp.Name, emp.Vacation)
			if err != nil {
				return fmt.Errorf("error writing data to file: %v", err)
			}
		}
	default:
		return fmt.Errorf("unsupported data type")
	}
	return nil
}
