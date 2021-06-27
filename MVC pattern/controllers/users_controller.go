package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../services"
	"../utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("id")
	userId, err := strconv.ParseInt(userIdParam, 10, 56)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		errorJson, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		_, _ = w.Write(errorJson)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		w.WriteHeader(apiErr.StatusCode)
		_, _ = w.Write([]byte(apiErr.Message))
		return
	}
	jsonValue, _ := json.Marshal(user)
	_, err = w.Write(jsonValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
