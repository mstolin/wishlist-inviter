package scrapper

type ScrapperFacade struct {
	AmazonScrapper string
}

func NewScrapperFacade(amazonScrapper string) ScrapperFacade {
	facade := ScrapperFacade{}
	facade.AmazonScrapper = amazonScrapper
	return facade
}
