package scrapper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/scrapper-facade/models"
)

func (facade ScrapperFacade) ScrapAmazonWishlist(whishlistId string) (*models.Whishlist, error) {
	whishlist := &models.Whishlist{}

	whishlistUrl := facade.Scrapper.AmazonScrapper + "/wishlist/" + whishlistId
	resp, err := http.Get(whishlistUrl)
	if err != nil {
		return whishlist, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return whishlist, err
	}

	if err := json.Unmarshal(body, whishlist); err != nil {
		return whishlist, err
	}

	return whishlist, nil
}
