package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/wishlist-service/models"
)

type ScrapperClient struct {
	URL string
}

func NewScrapperClient(url string) (ScrapperClient, error) {
	client := ScrapperClient{}

	if url == "" {
		return client, fmt.Errorf("SCRAPPER_URL is undefined")
	} else {
		client.URL = url
		return client, nil
	}
}

func (client ScrapperClient) GetWishlist(createReq *models.CreateReq) (*models.Whishlist, error) {
	whishlist := &models.Whishlist{}

	url := client.URL + "/" + createReq.Vendor + "/" + createReq.WishlistID
	resp, err := http.Get(url)
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
