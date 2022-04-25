package models

import (
	"fmt"
	"net/http"
)

type Mail struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

func (mail *Mail) Bind(request *http.Request) error {
	if mail.Recipient == "" || mail.Subject == "" || mail.Message == "" {
		return fmt.Errorf("recipient, subject, and message can't be empty")
	}
	return nil
}

func (mail *Mail) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
