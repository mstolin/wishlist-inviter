package models

import (
	"net/http"
)

type UserVerification struct {
	IsVerified bool `json:"is_verified"`
}

func (uv *UserVerification) Bind(request *http.Request) error {
	return nil
}

func (uv *UserVerification) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
