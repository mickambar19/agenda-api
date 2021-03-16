package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {

	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserFound(t *testing.T) {

	user, err := GetUser(123)

	assert.Nil(t, err, "we were not expecting an error")
	assert.NotNil(t, user)

	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Alex", user.FirstName)
	assert.EqualValues(t, "Jimenez", user.LastName)
	assert.EqualValues(t, "alex@yopmail.com", user.Email)
}
