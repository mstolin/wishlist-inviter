package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/wishlist-inviter/user-service/clients"
	"github.com/mstolin/wishlist-inviter/user-service/handler"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_ADAPTER")
	dbClient, err := clients.NewDatabaseClient(databaseUrl)
	if err != nil {
		log.Fatalf("Could not create Database Client: %v\n", err)
	}

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	signKey, exists := os.LookupEnv("JWT_SIGN_KEY")
	if !exists {
		log.Fatalf("No sign key given\n")
	}
	httpHandler := handler.NewHandler(signKey, dbClient)
	http.ListenAndServe(address, httpHandler)
}
