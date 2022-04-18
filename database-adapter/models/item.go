package models

import (
	"fmt"
	"net/http"
	"time"
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

type ItemList struct {
	Items []Item `json:"items"`
}

func (req *ItemList) Bind(r *http.Request) error {
	if len(req.Items) > 0 {
		return nil
	} else {
		return fmt.Errorf("items is empty")
	}
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
