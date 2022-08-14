package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/gmail-adapter/handler"
	"github.com/mstolin/present-roulette/gmail-adapter/mail"
)

func main() {
	host, port, sender, password := os.Getenv("GMAIL_HOST"), os.Getenv("GMAIL_PORT"), os.Getenv("GMAIL_MAIL"), os.Getenv("GMAIL_PASSWORD")
	smtpClient, error := mail.NewSMTPClient(host, port, sender, password)
	if error != nil {
		log.Fatalf("Could not create SMTP client: %v", error)
	}

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	signKey, exists := os.LookupEnv("JWT_SIGN_KEY")
	if !exists {
		log.Fatalf("No sign key given\n")
	}
	httpHandler := handler.NewHandler(signKey, smtpClient)
	if err := http.ListenAndServe(address, httpHandler); err != nil {
		log.Fatalf("Could not start HTTP server on address %s, %v", address, err)
	}
}
