package main

import "errors"

func TopThree(products []Product) ([3]Product, error) {
	if len(products) < 3 {

		return [3]Product{}, errors.New("not enough products: need 3")
	}
	return [3]Product(products[:3]), nil
}

func FromArray(arr [3]Product) []Product {
	s := arr[:]
	return s
}

func ApplyDiscount(products []Product, percent int) []Product {
	dp := make([]Product, 0, len(products))
	for _, p := range products {
		p.Price = p.Price * (100 - percent) / 100
		dp = append(dp, p)
	}
	return dp
}

func StockLevel(qty int) string {
	s := ""
	switch {
	case qty == 0:
		s = "out of stock"
	case qty > 0 && qty <= 5:
		s = "low stock"
	case qty > 5 && qty <= 20:
		s = "in stock"
	case qty > 20:
		s = "overstocked"
	}
	return s
}

func TotalValue(products ...Product) (total int, count int) {
	count = len(products)
	for _, p := range products {
		total += p.Price * p.Quantity
	}
	return
}

func FilterProducts(product []Product, predicate func(Product) bool) []Product {
	f := make([]Product, 0, len(product))
	for _, p := range product {
		if predicate(p) {
			f = append(f, p)
		}
	}
	return f
}
