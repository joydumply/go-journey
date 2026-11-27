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

	PrintStatus(&b)
	if err := b.CheckOut(); err != nil {
		fmt.Println("error: ", err)
	}
	PrintStatus(&b)

	if err := b.CheckOut(); err != nil {
		fmt.Println("error: ", err)
	}
	PrintStatus(&b)

	m1 := Magazine{
		Title:       "Nikita",
		IssueNumber: 777,
		Available:   true,
	}

	b.Return()
	c := []Item{&b, &m1}
	if err := b.CheckOut(); err != nil {
		fmt.Println("error: ", err)
	}

	PrintCatalog(c)

	q := Queue[string]{}
	q.Enqueue("Nik")
	q.Enqueue("Vlad")
	q.Enqueue("Vitos")

	for {
		name, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Println(name)
	}
	fmt.Println("queue is empty, len:", q.Len())

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

	d := CountByDecade(books)
	fmt.Println(d)

	for _, b := range books {
		y := AgeCategory(b.Year)
		fmt.Printf("%q is %s\n", b.Title, y)
	}

	t, a := LibraryStats(books...)
	fmt.Printf("total: %d, available: %d\n", t, a)

	fmt.Println(Describe(b))
	fmt.Println(Describe(&b))
	fmt.Println(Describe(m1))
	fmt.Println(Describe(42))

}
