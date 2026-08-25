package main

type Inventory map[string]Product

func (inv Inventory) Add(p Product) {
	inv[p.SKU] = p
}

func (inv Inventory) FindBySKU(sku string) (Product, bool) {
	p, ok := inv[sku]
	return p, ok
}

func (inv Inventory) RemoveBySKU(sku string) bool {
	if _, ok := inv[sku]; !ok {
		return false
	}
	delete(inv, sku)
	return true
}

func (inv Inventory) ToSlice() []Product {
	products := make([]Product, 0, len(inv))
	for _, p := range inv {
		products = append(products, p)
	}
	return products
}