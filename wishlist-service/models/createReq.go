package models

import (
	"errors"
	"net/http"
)

type CreateReq struct {
	WishlistID string `json:"wishlist_id"`
	Vendor     string `json:"vendor"`
	UserId     string `json:"user_id"`
}

func (createReq *CreateReq) Bind(request *http.Request) error {
	if createReq.WishlistID == "" || createReq.Vendor == "" || createReq.UserId == "" {
		return errors.New("wishlist_id, vendor, and user_id are required fields")
	}
	return nil
}
