package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mstolin/present-roulette/user-service/clients"
	"github.com/mstolin/present-roulette/user-service/handler"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_ADAPTER")
	dbClient, err := clients.NewDatabaseClient(databaseUrl)
	if err != nil {
		log.Fatalf("Could not create Database Client: %v", err)
	}

	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	httpHandler := handler.NewHandler("SUPER_SECRET", dbClient)
	http.ListenAndServe(address, httpHandler)
}
