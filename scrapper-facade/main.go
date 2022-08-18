package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/wishlist-inviter/scrapper-facade/handler"
	"github.com/mstolin/wishlist-inviter/scrapper-facade/scrapper"
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
	signKey, exists := os.LookupEnv("JWT_SIGN_KEY")
	if !exists {
		log.Fatalf("No sign key given\n")
	}
	httpHandler := handler.NewHandler(signKey, scrapperFacade)
	http.ListenAndServe(address, httpHandler)
}
