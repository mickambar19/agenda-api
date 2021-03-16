package services

import (
	"github.com/amjimenez/agenda-api/models"
	"github.com/amjimenez/agenda-api/utils"
)

func GetUser(userId int64) (*models.User, *utils.ApplicationError) {

	return models.GetUser(userId)
}
