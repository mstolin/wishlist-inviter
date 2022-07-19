package models

import (
	"fmt"
	"net/http"
)

type Mail struct {
	Recipient string `json:"recipient"`
	Body      string `json:"body"`
}

func (mail *Mail) Bind(request *http.Request) error {
	if mail.Recipient == "" || mail.Body == "" {
		return fmt.Errorf("a recipient, body is required")
	}
	return nil
}

func (mail *Mail) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
