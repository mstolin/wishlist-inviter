package models

import (
	"net/http"
)

type GMailResponse struct {
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
}

func (gmailResp *GMailResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
