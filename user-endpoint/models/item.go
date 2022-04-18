package models

type Item struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Vendor string  `json:"vendor"`
}
