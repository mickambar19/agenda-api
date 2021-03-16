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

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
