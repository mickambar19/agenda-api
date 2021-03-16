package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amjimenez/agenda-api/services"
	"github.com/amjimenez/agenda-api/utils"
)

// GetUsers retrieve the user passed in query
func GetUsers(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
