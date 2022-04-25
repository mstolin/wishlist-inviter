package models

import (
	"errors"
	"fmt"
	"net/http"
)

// The wishlist model is only used for scraping
type Whishlist struct {
	ID     string         `json:"id"`
	Vendor string         `json:"vendor"`
	Name   string         `json:"name"`
	Items  []WishlistItem `json:"items"`
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

// This wishlist item is only used for scraping
type WishlistItem struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Vendor string  `json:"vendor"`
}

func (item *WishlistItem) Bind(r *http.Request) error {
	if item.ID != "" || item.Name != "" || item.Price < 0 || item.Vendor != "" {
		return nil
	} else {
		return fmt.Errorf("id, name, price, and vendor can't be empty")
	}
}
