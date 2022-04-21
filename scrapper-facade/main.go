package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/scrapper-facade/handler"
	"github.com/mstolin/present-roulette/scrapper-facade/scrapper"
)

func main() {
	amazonScrapper := os.Getenv("AMAZON_SCRAPPER")
	if amazonScrapper == "" {
		log.Default().Println("AMAZON_SCRAPPER is empty")
	}
	scrapperFacade := scrapper.NewScrapperFacade(amazonScrapper)

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	httpHandler := handler.NewHandler(scrapperFacade)
	http.ListenAndServe(address, httpHandler)
}
