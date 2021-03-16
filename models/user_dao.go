package models

import (
	"fmt"
	"net/http"

	"github.com/amjimenez/agenda-api/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Alex", LastName: "Jimenez", Email: "alex@yopmail.com"},
	}
)

// GetUser retrieve data of specified user by id
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
