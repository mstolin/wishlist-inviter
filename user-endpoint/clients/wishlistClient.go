package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/models"
)

type WishlistClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Instantiates a new instance
func NewWishlistClient(url string) (WishlistClient, error) {
	client := WishlistClient{}
	if url == "" {
		return client, fmt.Errorf("wishlist service URL is not defined or empty")
	} else {
		client.URL = url
	}

	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends a request to the item service to import all items from an Amazon wishlist
func (client WishlistClient) ImportAmazonWishlist(wishlistId string) (models.ItemList, error) {
	itemLst := models.ItemList{}

	url := fmt.Sprintf("%s/amazon/wishlist/%s", client.URL, wishlistId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return itemLst, err
	}

	if err := json.Unmarshal(res, &itemLst); err != nil {
		return itemLst, err
	}
	return itemLst, nil
}
