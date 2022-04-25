package scrapper

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/models"
)

func (facade ScrapperFacade) ScrapAmazonWishlist(wishlistId string) (models.Whishlist, error) {
	wishlist := models.Whishlist{}
	if facade.AmazonScrapper == "" {
		return wishlist, fmt.Errorf("can't establish connection to Amazon Scrapper, because AMAZON_SCRAPPER is empty")
	}

	url := fmt.Sprintf("%s/wishlist/%s", facade.AmazonScrapper, wishlistId) // TODO wishlist => wishlists
	res, err := httpFacadeInstance.DoGet(url)
	if err != nil {
		return wishlist, err
	}

	if err := json.Unmarshal(res, &wishlist); err != nil {
		return wishlist, err
	}

	return wishlist, nil
}
