package scrapper

import (
	"fmt"
	"os"
)

type ScrapperFacade struct {
	Scrapper scrapper
}

type scrapper struct {
	AmazonScrapper string
}

func getScrapper() scrapper {
	scrapper := scrapper{}

	if amazonScrapper := os.Getenv("AMAZON_SCRAPPER"); amazonScrapper != "" {
		scrapper.AmazonScrapper = amazonScrapper
	} else {
		fmt.Fprint(os.Stdout, "AMAZON_SCRAPPER is undefined")
	}

	return scrapper
}

func NewScrapperFacade() ScrapperFacade {
	facade := ScrapperFacade{}
	facade.Scrapper = getScrapper()
	return facade
}
