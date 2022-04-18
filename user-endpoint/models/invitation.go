package models

import (
	"errors"
	"net/http"
)

type Invitation struct {
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
}

func (invitation *Invitation) Bind(request *http.Request) error {
	if invitation.Subject == "" || invitation.Message == "" || invitation.Recipient == "" {
		return errors.New("subject, message, and recipient are required fields")
	}
	return nil
}

func (Invitation *Invitation) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
