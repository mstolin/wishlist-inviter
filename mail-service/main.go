package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/mail-service/database"
	"github.com/mstolin/present-roulette/mail-service/gmail"
	"github.com/mstolin/present-roulette/mail-service/handler"
	"github.com/mstolin/present-roulette/mail-service/messages"
)

func main() {
	// GMail client
	gmailService := os.Getenv("GMAIL_ADAPTER")
	gmailClient, err := gmail.NewGMailClient(gmailService)
	if err != nil {
		log.Fatalf("Could not create GMail client: %v", err)
	}
	// DB client
	dbService := os.Getenv("DATABASE_ADAPTER")
	dbClient, err := database.NewDatabaseClient(dbService)
	if err != nil {
		log.Fatalf("Could not create Database Client: %v", err)
	}
	// Message factory
	msgFactory := messages.MessageFactory{}

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	httpHandler := handler.NewHandler(gmailClient, dbClient, msgFactory)
	http.ListenAndServe(address, httpHandler)
}
