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

	"github.com/mstolin/present-roulette/database-adapter/db"
	"github.com/mstolin/present-roulette/database-adapter/handler"
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

	dbHost, dbPort, dbUser, dbPassword, dbName :=
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	database, err := db.Initialize(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not initialize database connection at %s:***@%s:%s/%s, %v",
			dbUser, dbHost, dbPort, dbName, err)
	} else {
		log.Default().Printf("Successfully connected to %s:***@%s/%s:%s",
			dbUser, dbHost, dbPort, dbName)
	}

	httpHandler := handler.NewHandler(database)
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
