package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := Book{
		Author:    Author{FirstName: "Nikita", LastName: "Babichev", Country: "Ukraine"},
		Title:     "Book one",
		ISBN:      "1q2w3erty",
		Year:      2026,
		Available: true,
		Notes:     "Test note",
	}

	fmt.Println(b.FirstName)

	data, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(data))

	// PrintStatus(&b)
	// if err := b.CheckOut(); err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// PrintStatus(&b)

	// if err := b.CheckOut(); err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// PrintStatus(&b)

	m1 := Magazine{
		Title:       "Nikita",
		IssueNumber: 777,
		Available:   true,
	}

	// d1 := b.Describe()
	// d2 := m1.Describe()
	// fmt.Println(d1)
	// fmt.Println(d2)

	c := []Item{&b, &m1}
	b.CheckOut()

	PrintCatalog(c)

}
