package main

import (
	"crm/internal/data"
	"crm/internal/models"
	"crm/pkg/dbutil"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current dir: %v", err)
	}
	ur, err := models.NewUserRepository(filepath.Join(dir, "storage"))
	if err != nil {
		log.Printf("models error:\n%v", err)
	}
	err = dbutil.ExecSQLFile(ur.Db, filepath.Join("sql", "users.sql"))
	if err != nil {
		log.Printf("dbutil error:\n%v", err)
	}

	usersTXT, err := data.ParseUsersTXT(filepath.Join("storage", "users.txt"))
	if err != nil {
		log.Printf("error parsing users.txt:\n%v", err)
	}
	users, err := data.CompleteUsersTXT(usersTXT)
	if err != nil {
		log.Fatalf("error completing users:\n%v", err)
	}
	err = ur.CreateUsers(users)
	if err != nil {
		log.Fatalf("error creating users in db:\n%v", err)
	}
	usersDb, err := ur.GetAllUsers()
	for _, user := range usersDb {
		fmt.Printf("%+v\n", user)
	}
	err = ur.Close()
	if err != nil {
		log.Fatalf("error closing db")
	}
}
