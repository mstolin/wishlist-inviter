package models

import (
	"fmt"
	"net/http"
)

type SuccessMessage struct {
	Message string `json:"message"`
}

func (m *SuccessMessage) Bind(request *http.Request) error {
	if m.Message == "" {
		return fmt.Errorf("message is required")
	}
	return nil
}

func (m *SuccessMessage) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
