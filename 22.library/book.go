package main

import (
	"errors"
	"fmt"
)

type Item interface {
	Describe() string
}

type Borrowable interface {
	CheckOut() error
	Return()
}

// LibraryItem = Item + Borrowable
// Magazine не удовлетворяет LibraryItem т.к в нем не реализованы методы Checkout и Return
// Book (значение) тоже не удовлетворяет,т.к. в Book методы сделаны с ссылочным значением
// *Book удовлетворяет, потому что описанные методы реализованы с ссылкой, а не с копией 
type LibraryItem interface {
	Item
	Borrowable
}
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
}

type Book struct {
	Author    `json:"author"`
	Title     string `json:"title"`
	ISBN      string `json:"isbn"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
	Notes     string `json:"notes,omitempty"`
}

type Magazine struct {
	Title       string `json:"title"`
	IssueNumber int    `json:"issueNumber"`
	Available   bool   `json:"available"`
}

func (b *Book) CheckOut() error {
	if !b.Available {
		return errors.New("book is not available")
	}
	b.Available = false
	return nil
}

func (b *Book) Return() {
	b.Available = true
}

func PrintStatus(b *Book) {
	if b.Available {
		fmt.Printf("%q is available\n", b.Title)
		return
	}
	fmt.Printf("%q is not available\n", b.Title)
}

func (b Book) Describe() string {
	var available string
	if b.Available {
		available = "available"
	} else {
		available = "not available"
	}
	return fmt.Sprintf("Book: %q by %s %s (%d), %s", b.Title, b.FirstName, b.LastName, b.Year, available)
}

func (m Magazine) Describe() string {
	var available string
	if m.Available {
		available = "available"
	} else {
		available = "not available"
	}
	return fmt.Sprintf("Magazine: %q #%d, %s", m.Title, m.IssueNumber, available)
}

func PrintCatalog(items []Item) {
	for _, v := range items {
		fmt.Println(v.Describe())
	}
}

func Describe(x any) string {
	switch x.(type) {
	case Book, *Book:
		return "it's a book"
	case Magazine, *Magazine:
		return "it's a magazine"
	default:
		return "unknown item"
	}
}
