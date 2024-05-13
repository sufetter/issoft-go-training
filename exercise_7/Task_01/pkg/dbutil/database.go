package dbutil

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func Create(path string, name string) (*os.File, error) {
	fmt.Println(filepath.Join(path, name+".db"))
	db, err := os.Create(filepath.Join(path, name+".db"))
	if err != nil {
		return nil, fmt.Errorf("error creating database: %w", err)
	}
	return db, nil
}
func Open(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error сonnecting database: %w", err)
	}

	return db, nil
}

//Подумал, что вдруг кому-то надо будет создать или
//открыть бд отдельно, а это будет чтоб сразу все

func CreateAndOpen(path string, name string) (*sql.DB, error) {
	_, err := Create(path, name)
	if err != nil {
		return nil, err
	}

	db, err := Open(filepath.Join(path, name+".db"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExecSQLFile(db *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query != "" {
			_, err := db.Exec(query)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
