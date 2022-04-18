package main

import (
	"net/http"
	"os"
)

func main() {
	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}

	//httpHandler =

	wishlistService := os.Getenv("WISHLIST_SERVICE")
	mailService := os.Getenv("MAIL_SERVICE")
	userService := os.Getenv("USER_SERVICE")

	http.ListenAndServe(address, nil)
}
