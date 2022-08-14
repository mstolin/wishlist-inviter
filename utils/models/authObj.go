package models

import (
	"errors"
	"net/http"
)

type AuthObj struct {
	UserId string `json:"user_id"`
}

func (ao *AuthObj) Bind(request *http.Request) error {
	if ao.UserId == "" {
		return errors.New("user ID is required")
	}
	return nil
}

func (ao *AuthObj) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
