package scrapper

import "github.com/mstolin/wishlist-inviter/utils/clients"

type ScrapperFacade struct {
	AmazonScrapper string
}

var httpFacadeInstance = clients.NewHTTPFacade()

func NewScrapperFacade(amazonScrapper string) ScrapperFacade {
	facade := ScrapperFacade{}
	facade.AmazonScrapper = amazonScrapper

	return facade
}
