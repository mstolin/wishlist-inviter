package clients

import (
	"encoding/json"
	"fmt"

	"github.com/mstolin/present-roulette/utils/clients"
	"github.com/mstolin/present-roulette/utils/httpErrors"
	"github.com/mstolin/present-roulette/utils/models"
)

type UserClient struct {
	URL        string
	httpFacade clients.HTTPFacade
}

// Constructs a new user client instance.
func NewUserClient(url string) (UserClient, error) {
	client := UserClient{}
	if url == "" {
		return client, fmt.Errorf("user service URL is not defined or empty")
	} else {
		client.URL = url
	}

	client.httpFacade = clients.NewHTTPFacade()
	return client, nil
}

// Sends a request to the user server to create an empty user.
func (client UserClient) CreateEmptyUser() (models.User, *httpErrors.ErrorResponse) {
	user := models.User{}
	url := fmt.Sprintf("%s/users", client.URL)

	jsonStr, err := json.Marshal("{}")
	if err != nil {
		return user, httpErrors.ErrBadRequestRenderer(err)
	}
	res, httpErr := client.httpFacade.DoPost(url, jsonStr)
	if httpErr != nil {
		return user, httpErr
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, httpErrors.ErrBadRequestRenderer(err)
	}
	return user, nil
}

// Sends a GET request to the user service to receive the user given the ID.
func (client UserClient) GetUser(userId string) (models.User, *httpErrors.ErrorResponse) {
	user := models.User{}
	url := fmt.Sprintf("%s/users/%s", client.URL, userId)

	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, httpErrors.ErrBadRequestRenderer(err)
	}
	return user, nil
}

// Returns all items of a specific user.
func (client UserClient) GetUserItems(userId string) (models.ItemList, *httpErrors.ErrorResponse) {
	itemLst := models.ItemList{}
	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)

	res, err := client.httpFacade.DoGet(url)
	if err != nil {
		return itemLst, err
	}

	if err := json.Unmarshal(res, &itemLst); err != nil {
		return itemLst, httpErrors.ErrBadRequestRenderer(err)
	}
	return itemLst, nil
}

// Adds an item list to a specific user.
func (client UserClient) AddUserItems(userId string, itemLst []models.Item) (models.ItemList, *httpErrors.ErrorResponse) {
	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)

	jsonStr, err := json.Marshal(itemLst)
	if err != nil {
		return itemLst, httpErrors.ErrBadRequestRenderer(err)
	}

	res, httpErr := client.httpFacade.DoPost(url, jsonStr)
	if httpErr != nil {
		return itemLst, httpErr
	}

	if err := json.Unmarshal(res, &itemLst); err != nil {
		return itemLst, httpErrors.ErrBadRequestRenderer(err)
	}
	return itemLst, nil
}

// Updates an item based on the givne update model
func (client UserClient) UpdateItem(userId string, itemId int, update models.ItemUpdate) (models.Item, *httpErrors.ErrorResponse) {
	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	var item models.Item

	jsonStr, err := json.Marshal(update)
	if err != nil {
		return item, httpErrors.ErrBadRequestRenderer(err)
	}

	res, httpErr := client.httpFacade.DoPut(url, jsonStr)
	if httpErr != nil {
		return item, httpErr
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return item, httpErrors.ErrBadRequestRenderer(err)
	}
	return item, nil
}
