package models

import (
	"errors"
	"net/http"
)

type Invitation struct {
	UserId    string `json:"user_id"`
	Items     []uint `json:"items"`
	Recipient string `json:"recipient"`
}

func (invitation *Invitation) Bind(request *http.Request) error {
	if invitation.Recipient == "" || invitation.UserId == "" || len(invitation.Items) <= 0 {
		return errors.New("user_id, items, and recipient are required fields")
	}
	return nil
}

func (Invitation *Invitation) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
