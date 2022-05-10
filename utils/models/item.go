package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Item struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(255); not null" json:"name"`
	Price     float32   `sql:"type:decimal(10,2); not null" json:"price"`                   // TODO better to rid of this in the future
	Vendor    string    `gorm:"type:varchar(100); not null" json:"vendor"`                  // TODO future table called Vendor, foreign key here
	VendorID  string    `gorm:"type:varchar(255); unique_index; not null" json:"vendor_id"` // ID given by the vendor TODO better name
}

func (*Item) TableName() string {
	return "item"
}

func (i *Item) Bind(r *http.Request) error {
	if i.Name == "" || i.Vendor == "" || i.Price <= 0 {
		return fmt.Errorf("It seems like some properties are undefined or incorrect (Name: '%s', Vendor: '%s', Price: '%d')",
			i.Name, i.Vendor, i.Price)
	}
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ItemList []Item

func (*ItemList) Bind(r *http.Request) error {
	return nil
}

type ItemResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Vendor    string    `json:"vendor"`
	VendorID  string    `json:"vendor_id"`
}

func NewItemReponse(item Item) ItemResponse {
	res := ItemResponse{}
	res.ID = item.ID
	res.CreatedAt = item.CreatedAt
	res.UpdatedAt = item.UpdatedAt
	res.Name = item.Name
	res.Price = item.Price
	res.Vendor = item.Vendor
	res.VendorID = item.VendorID
	return res
}

func (*ItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ItemResponseList []ItemResponse

func NewItemResponseListRenderer(items ItemList) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		res := NewItemReponse(item)
		list = append(list, &res)
	}
	return list
}
