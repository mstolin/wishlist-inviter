package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mstolin/present-roulette/user-service/models"
)

type DatabaseClient struct {
	URL string
}

func NewDatabaseClient(url string) (DatabaseClient, error) {
	client := DatabaseClient{}

	if url == "" {
		return client, fmt.Errorf("DATABASE_URL is undefined")
	} else {
		client.URL = url
		return client, nil
	}
}

func (client DatabaseClient) CreateUser() (*models.User, error) {
	response := &models.User{}

	jsonStr, err := json.Marshal("{\"items\": []}") // empty user with no items
	if err != nil {
		return response, err
	}

	url := client.URL + "/users"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if err := json.Unmarshal(body, response); err != nil {
		return response, err
	}

	return response, nil
}

func (client DatabaseClient) GetUser(userId string) (models.User, error) {
	user := models.User{}

	url := client.URL + "/users"
	resp, err := http.Get(url)
	if err != nil {
		return user, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(body, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (client DatabaseClient) DeleteUser(userId string) (models.User, error) {
	user := models.User{}

	url := client.URL + "/users/" + userId
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return user, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(body, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (client DatabaseClient) GetItemsByUser(userId string) (models.ItemList, error) {
	list := models.ItemList{}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	resp, err := http.Get(url)
	if err != nil {
		return list, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}
	if err := json.Unmarshal(body, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (client DatabaseClient) AddItemsToUser(userId string, items models.ItemList) (models.ItemList, error) {
	addedItems := models.ItemList{}

	jsonStr, err := json.Marshal(items)
	if err != nil {
		return addedItems, err
	}

	url := fmt.Sprintf("%s/users/%s/items", client.URL, userId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return addedItems, err
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return addedItems, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return addedItems, err
	}
	if err := json.Unmarshal(body, &addedItems); err != nil {
		return addedItems, err
	}

	return addedItems, nil
}

func (client DatabaseClient) GetItemByUser(userId string, itemId int) (models.Item, error) {
	item := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	resp, err := http.Get(url)
	if err != nil {
		return item, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}
	if err := json.Unmarshal(body, &item); err != nil {
		return item, err
	}

	return item, nil
}

func (client DatabaseClient) UpdateItemByUser(userId string, itemId int, update models.Item) (models.Item, error) {
	updatedItem := models.Item{}

	jsonStr, err := json.Marshal(update)
	if err != nil {
		return updatedItem, err
	}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return updatedItem, err
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return updatedItem, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return updatedItem, err
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		return updatedItem, err
	}

	return updatedItem, nil
}

func (client DatabaseClient) DeleteItemByUser(userId string, itemId int) (models.Item, error) {
	deletedItem := models.Item{}

	url := fmt.Sprintf("%s/users/%s/items/%d", client.URL, userId, itemId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return deletedItem, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return deletedItem, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deletedItem, err
	}
	if err := json.Unmarshal(body, &deletedItem); err != nil {
		return deletedItem, err
	}

	return deletedItem, nil
}
