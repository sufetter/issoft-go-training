package models

import (
	"database/sql"
	"log"
	"webcrm/pkg/dbutil"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(dbPath string) (*UserRepository, error) {
	db, err := dbutil.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &UserRepository{db}, nil
}

type User struct {
	Id       int
	Email    string
	Name     string
	IsActive bool
}

func (ur *UserRepository) GetUserById(id int) (*User, error) {
	var user User
	err := ur.Db.QueryRow("SELECT Id, Email, Name, IsActive FROM Users WHERE Id = ?", id).Scan(&user.Id, &user.Email, &user.Name, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]*User, error) {
	rows, err := ur.Db.Query("SELECT Id, Email, Name, IsActive FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.Db.Exec("DELETE FROM Users WHERE Id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Close() error {
	err := ur.Db.Close()
	if err != nil {
		log.Printf("Error closing database: %v", err)
		return err
	}
	return nil
}
