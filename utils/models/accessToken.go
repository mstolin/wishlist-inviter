package models

import (
	"errors"
	"net/http"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

func (at *AccessToken) Bind(request *http.Request) error {
	if at.AccessToken == "" {
		return errors.New("access token is required")
	}
	return nil
}

func (at *AccessToken) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
