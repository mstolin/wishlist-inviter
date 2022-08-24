package scrapper

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/wishlist-inviter/utils/httpErrors"
	"github.com/mstolin/wishlist-inviter/utils/models"
)

func (facade ScrapperFacade) ScrapAmazonWishlist(wishlistId, accessToken string) (models.Wishlist, *httpErrors.ErrorResponse) {
	wishlist := models.Wishlist{}
	if facade.AmazonScrapper == "" {
		err := fmt.Errorf("can't establish connection to Amazon Scrapper, because AMAZON_SCRAPPER is empty")
		return wishlist, httpErrors.ErrServerErrorRenderer(err)
	}

	url := fmt.Sprintf("%s/wishlists/%s", facade.AmazonScrapper, wishlistId)
	res, httpErr := httpFacadeInstance.DoGet(url, accessToken)
	if httpErr != nil {
		return wishlist, httpErr
	}

	if err := json.Unmarshal(res, &wishlist); err != nil {
		return wishlist, httpErrors.ErrServerErrorRenderer(err)
	}

	return wishlist, nil
}
