package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mstolin/present-roulette/wishlist-service/clients"
	"github.com/mstolin/present-roulette/wishlist-service/handler"
)

func main() {
	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	listener, error := net.Listen("tcp", address)
	if error != nil {
		log.Fatalf("Error occurred: %s", error.Error())
	}

	scrapperUrl := os.Getenv("SCRAPPER_URL")
	scrapperClient, err := clients.NewScrapperClient(scrapperUrl)
	if err != nil {
		log.Fatalf("Could not create Scrapper Client: %v", error)
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	dbClient, err := clients.NewDatabaseClient(databaseUrl)
	if err != nil {
		log.Fatalf("Could not create Database Client: %v", error)
	}

	httpHandler := handler.NewHandler(scrapperClient, dbClient)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", address)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
