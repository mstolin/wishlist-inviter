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
	if invitation.Subject == "" || invitation.Recipient == "" || invitation.Message == "" {
		return errors.New("subject, message, and recipient are required fields")
	}
	return nil
}

func (Invitation *Invitation) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type InvitationReq struct {
	Subject   string `json:"subject"`
	UserId    string `json:"user_id"`
	Items     []uint `json:"items"`
	Recipient string `json:"recipient"`
}

func (invitation *InvitationReq) Bind(request *http.Request) error {
	if invitation.Subject == "" || invitation.Recipient == "" || invitation.UserId == "" || len(invitation.Items) <= 0 {
		return errors.New("subject, user_id, items, and recipient are required fields")
	}
	return nil
}

func (Invitation *InvitationReq) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
