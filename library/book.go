package main

import (
	"errors"
	"fmt"
)

// Part 1

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

// Part 1 end

// Part 2

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
