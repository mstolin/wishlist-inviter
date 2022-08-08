package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/user-endpoint/clients"
	"github.com/mstolin/present-roulette/user-endpoint/handler"
)

func main() {
	userClient, mailClient, wishlistClient := initClients()

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	httpHandler := handler.NewHandler(userClient, mailClient, wishlistClient)
	http.ListenAndServe(address, httpHandler)
}

func initClients() (clients.UserClient, clients.MailClient, clients.ScrapperFacadeClient) {
	userServiceUrl := os.Getenv("USER_SERVICE")
	userClient, err := clients.NewUserClient(userServiceUrl)
	if err != nil {
		log.Fatalf("Could not init UserClient: %v", err)
	}

	mailServiceUrl := os.Getenv("MAIL_SERVICE")
	mailClient, err := clients.NewMailClient(mailServiceUrl)
	if err != nil {
		log.Fatalf("Could not init MailClient: %v", err)
	}

	wishlistServiceUrl := os.Getenv("SCRAPPER_FACADE")
	wishlistClient, err := clients.NewScrapperFacadeClient(wishlistServiceUrl)
	if err != nil {
		log.Fatalf("Could not init ScrapperFacade: %v", err)
	}

	return userClient, mailClient, wishlistClient
}
