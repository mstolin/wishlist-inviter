package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/mail-adapter/gmail"
	"github.com/mstolin/present-roulette/mail-adapter/handler"
)

func main() {
	gmailService := os.Getenv("GMAIL_SERVICE")
	gmailClient, error := gmail.NewGMailClient(gmailService)
	if error != nil {
		log.Fatalf("Could not create GMail client: %v", error)
	}

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	httpHandler := handler.NewHandler(gmailClient)
	http.ListenAndServe(address, httpHandler)
}
