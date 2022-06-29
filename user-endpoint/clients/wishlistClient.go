package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
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

/*
	Also das problem is folgendes: hier bekommen wir wishlist zurueck, aber wir brauchen nur die itemlist
	fur alles weitere, daher itemlist aus wishlist rausnehmen und nur die weitergeben ...
*/

// Sends a request to the item service to import all items from an Amazon wishlist
func (client WishlistClient) ImportAmazonWishlist(wishlistId string) (models.Wishlist, *httpErrors.ErrorResponse) {
	wishlist := models.Wishlist{}

	url := fmt.Sprintf("%s/amazon/wishlists/%s", client.URL, wishlistId)
	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return wishlist, err
	}

	if err := json.Unmarshal(res, &wishlist); err != nil {
		return wishlist, httpErrors.ErrBadRequestRenderer(err)
	}
	return wishlist, nil
}
