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

	q := Queue[string]{}
	q.Enqueue("Nik")
	q.Enqueue("Vlad")
	q.Enqueue("Vitos")

	item, ok := q.Dequeue()
	fmt.Println(item, ok, q.Len())

	books := []Book{
		{
			Author:    Author{FirstName: "Fyodor", LastName: "Dostoevsky", Country: "Russia"},
			Title:     "Crime and Punishment",
			ISBN:      "978-0143058144",
			Year:      1866,
			Available: true,
		},
		{
			Author:    Author{FirstName: "George", LastName: "Orwell", Country: "UK"},
			Title:     "1984",
			ISBN:      "978-0451524935",
			Year:      1949,
			Available: false,
		},
		{
			Author:    Author{FirstName: "Alan", LastName: "Donovan", Country: "USA"},
			Title:     "The Go Programming Language",
			ISBN:      "978-0134190440",
			Year:      2015,
			Available: true,
		},
		{
			Author:    Author{FirstName: "Mikhail", LastName: "Bulgakov", Country: "Russia"},
			Title:     "The Master and Margarita",
			ISBN:      "978-0141180144",
			Year:      1967,
			Available: false,
		},
		{
			Author:    Author{FirstName: "Taras", LastName: "Shevchenko", Country: "Ukraine"},
			Title:     "Kobzar",
			ISBN:      "978-9660353005",
			Year:      1840,
			Available: true,
			Notes:     "First edition reprint",
		},
	}

	available := Filter(books, func(b Book) bool {
		return b.Available
	})

	for _, b := range available {
		fmt.Println(b.Describe())
	}

}
