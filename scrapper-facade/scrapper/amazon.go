package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/scrapper-facade/models"
)

func (facade ScrapperFacade) ScrapAmazonWishlist(wishlistId string) (models.Whishlist, error) {
	wishlist := models.Whishlist{}
	if facade.AmazonScrapper == "" {
		return wishlist, fmt.Errorf("can't establish connection to Amazon Scrapper, because AMAZON_SCRAPPER is empty")
	}

	wishlistUrl := fmt.Sprintf("%s/wishlist/%s", facade.AmazonScrapper, wishlistId)
	resp, err := http.Get(wishlistUrl)
	if err != nil {
		return wishlist, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wishlist, err
	}

	if err := json.Unmarshal(body, &wishlist); err != nil {
		return wishlist, err
	}

	return wishlist, nil
}
