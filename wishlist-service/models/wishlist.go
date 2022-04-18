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

func (whishlist *Whishlist) Bind(r *http.Request) error {
	if whishlist.ID == "" || whishlist.Vendor == "" || whishlist.Name == "" {
		return errors.New("id, vendor, and name are required fields")
	}
	return nil
}

func (whishlist *Whishlist) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
