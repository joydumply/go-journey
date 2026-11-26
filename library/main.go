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
}
