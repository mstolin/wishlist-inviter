package models

import (
	"errors"
	"net/http"
)

type Whishlist struct {
	ID     string `json:"id"`
	Vendor string `json:"vendor"`
	Name   string `json:"name"`
	Items  []Item `json:"items"`
}

func (whishlist *Whishlist) Bind(request *http.Request) error {
	if whishlist.Name == "" {
		return errors.New("name is a required field")
	}
	return nil
}

func (whishlist *Whishlist) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
