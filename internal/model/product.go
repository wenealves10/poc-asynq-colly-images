package model

type Product struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Images []string `json:"images"`
}
