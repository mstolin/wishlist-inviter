package models

import (
	"errors"
	"net/http"
)

type InvitationReq struct {
	Subject   string `json:"subject"`
	Recipient string `json:"recipient"`
}

func (invitation *InvitationReq) Bind(request *http.Request) error {
	if invitation.Subject == "" || invitation.Recipient == "" {
		return errors.New("subject and recipient are required fields")
	}
	return nil
}

func (Invitation *InvitationReq) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type InvitationRes struct {
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
}

func (invitation *InvitationRes) Bind(request *http.Request) error {
	if invitation.Subject == "" || invitation.Recipient == "" {
		return errors.New("subject and recipient are required fields")
	}
	return nil
}

func (Invitation *InvitationRes) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
