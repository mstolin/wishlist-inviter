package clients

import "fmt"

type WishlistClient struct {
	URL string
}

func NewWishlistClient(url string) (WishlistClient, error) {
	client := WishlistClient{}
	if url == "" {
		return client, fmt.Errorf("wishlist service URL is not defined or empty")
	} else {
		client.URL = url
	}

	return client, nil
}

// Create a new wishlist
func (client WishlistClient) importWishlist() error {
	return nil
}

func (client WishlistClient) getWishlist(user string) error {
	return nil
}
