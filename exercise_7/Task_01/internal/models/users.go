package models

import (
	"crm/pkg/dbutil"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(dbPath string) (*UserRepository, error) {
	db, err := dbutil.CreateAndOpen(dbPath, "users")
	if err != nil {
		return nil, err
	}
	return &UserRepository{db}, nil
}

type User struct {
	Id       int
	Email    string
	password string
	Name     string
	IsActive bool
}

func NewUser(name, surname string) (*User, error) {
	email := GenerateUserEmail(name, surname)

	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password := string(hash)

	return &User{
		Email:    email,
		password: password,
		Name:     name + " " + surname,
		IsActive: true,
	}, nil
}
func GenerateUserEmail(name string, surname string) string {
	name = strings.ToLower(name)
	surname = strings.ToLower(surname)

	name = strings.ReplaceAll(name, " ", "")
	surname = strings.ReplaceAll(surname, " ", "")

	name = FirstRuneToUpper(name)
	surname = FirstRuneToUpper(surname)

	email := fmt.Sprintf("%s%s@coolcompany.com", name, surname)

	return email
}

//можно было strings.Title, но
//этот метод устаревший, а использовать
//сторонний пакет я не нашел необходимости

func FirstRuneToUpper(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

func (ur *UserRepository) CreateUsers(users []*User) error {
	query, err := ur.Db.Prepare("INSERT INTO Users (Email, PasswordHash, Name, IsActive) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("eror preparing query: %v", err)
		return err
	}
	defer query.Close()

	for _, user := range users {
		_, err = query.Exec(user.Email, user.password, user.Name, user.IsActive)
		if err != nil {
			log.Printf("eror creating user: %v", err)
			return err
		}
	}

	return nil
}

func (ur *UserRepository) GetAllUsers() ([]*User, error) {
	rows, err := ur.Db.Query("SELECT Id, Email, PasswordHash, Name, IsActive FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Пароль выводим, тк указано вывести ВСЕ

	var users []*User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.password, &user.Name, &user.IsActive)
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

func (ur *UserRepository) Close() error {
	err := ur.Db.Close()
	if err != nil {
		log.Printf("Error closing database: %v", err)
		return err
	}
	return nil
}
