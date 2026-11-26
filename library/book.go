package main

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
