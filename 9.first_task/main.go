package main

import "fmt"

func main() {
	inv := make(Inventory)
	p1 := Product{
		ID:          1,
		Name:        "iphone",
		SKU:         "ABC123",
		Price:       1000,
		Quantity:    10,
		Description: "This is iPhone. All you need to know",
		Dimensions: Dimensions{
			Width:  10,
			Height: 10,
			Depth:  10,
		},
	}
	inv.Add(p1)

	// fmt.Println(inv.RemoveBySKU("ABC123"))

	inv.Add(Product{SKU: "A", Name: "Phone"})
	inv.Add(Product{SKU: "B", Name: "Laptop"})
	inv.Add(Product{SKU: "C", Name: "Flashlight"})

	products := inv.ToSlice()
	// fmt.Println(products)
	top, err := TopThree(products)
	fmt.Println(top, err)
	fmt.Println("=====")

	s := FromArray(top)
	// s[0] = Product{SKU: "CHANGED"}
	fmt.Println(top[0].SKU)
	fmt.Println(s[0].SKU)
	fmt.Println("=====")
	fmt.Println("DISCOUNTS!!!")
	original := []Product{{Price: 1000}, {Price: 2000}}
	discounted := ApplyDiscount(original, 10)

	fmt.Println(original)
	fmt.Println(discounted)

	fmt.Println(s)
	fmt.Println(StockLevel(s[0].Quantity))
	fmt.Println("=====")
	pr1 := Product{Price: 1000, Quantity: 2}
	pr2 := Product{Price: 500, Quantity: 3}
	prs := []Product{pr1, pr2}
	total, count := TotalValue(prs...)
	fmt.Println(total, count)

	unfilteredProducts := []Product{
		{SKU: "A", Quantity: 2},
		{SKU: "B", Quantity: 10},
		{SKU: "C", Quantity: 1},
	}

	lowStock := FilterProducts(unfilteredProducts, func(p Product) bool {
		return p.Quantity < 5
	})

	fmt.Println("=========")
	fmt.Println("Low Stock")
	fmt.Println(lowStock)

}
