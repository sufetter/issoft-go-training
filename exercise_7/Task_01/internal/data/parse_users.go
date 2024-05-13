package data

import (
	"crm/internal/models"
	"crm/pkg/parser"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type UserTXT struct {
	name    string
	surname string
}

func ParseUsersTXT(path string) ([]*UserTXT, error) {
	letterRegex := regexp.MustCompile(`^[a-zA-Z]+$`)

	parseFunc := func(line string) (*UserTXT, error) {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("expected 2 fields, found %d", len(fields))
		}
		for _, field := range fields {
			if !letterRegex.MatchString(field) {
				return nil, fmt.Errorf("non-letter characters found")
			}
		}
		return &UserTXT{
			name:    fields[0],
			surname: fields[1],
		}, nil
	}

	parsedUsers, err := parser.TXTFile(path, parseFunc)
	if err != nil {
		return nil, err
	}

	return parsedUsers, nil
}

func CompleteUsersTXT(usersTXT []*UserTXT) ([]*models.User, error) {
	users := make([]*models.User, 0, len(usersTXT))

	for _, user := range usersTXT {

		u, err := models.NewUser(user.name, user.surname)
		if err != nil {
			log.Printf("Error completing user: %v", err)
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}
