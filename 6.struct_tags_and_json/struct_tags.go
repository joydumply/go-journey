package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"` // will ignore in Marshal
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:",omitempty"`   // will left name as Name
	Price int    `json:"price,string"` // will parse to string
}

func main() {
	u := User{Name: "Anna", Age: 28, Email: "anna@example.com", Password: "Qwerty123"}
	u2 := User{Name: "Nik", Age: 29}
	data, _ := json.Marshal(u)
	d2, _ := json.Marshal(u2) // {"name":"Nik","age":29} no email
	fmt.Println(string(data))
	fmt.Println(string(d2))
	/*
		{"Name":"Anna","Age":28} without tags
		{"Name":"Anna","Age":28} with tags
	*/

	p1 := Product{ID: 1, Name: "Phone", Price: 1000}
	d3, _ := json.Marshal(p1)
	fmt.Println(string(d3)) // {"id":1,"Name":"Phone","price":"1000"}

}
