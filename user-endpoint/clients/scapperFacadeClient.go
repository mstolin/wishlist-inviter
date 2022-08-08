package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

type ScrapperFacadeClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Instantiates a new instance
func NewScrapperFacadeClient(url string) (ScrapperFacadeClient, error) {
	client := ScrapperFacadeClient{}
	if url == "" {
		return client, fmt.Errorf("scrapper facade URL is not defined or empty")
	} else {
		client.URL = url
	}

	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends a request to the item service to import all items from an Amazon wishlist
func (client ScrapperFacadeClient) ImportAmazonWishlist(wishlistId string) (models.Wishlist, *httpErrors.ErrorResponse) {
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
