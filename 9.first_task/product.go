package main

import (
	"encoding/json"
)

type Dimensions struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

type Product struct {
	ID          int    `json:"product_id"`
	Name        string `json:"name"`
	SKU         string `json:"sku"` // unique article
	Price       int    `json:"price,string"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description,omitempty"`
	Dimensions
}

func ParseProduct(data []byte) (Product, error) {
	var p Product
	err := json.Unmarshal(data, &p)

	return p, err
}
