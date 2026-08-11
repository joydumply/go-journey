package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Address struct {
	City    string
	Country string
}

type User struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

func main() {
	fmt.Println(Vertex{1, 2})

	u1 := User{Name: "Anna", Age: 28, Email: "anna@example.com"}

	fmt.Println(u1)

	fmt.Println("User1 Name: ", u1.Name)

	// Not a pointer so it's copying
	u2 := u1
	u2.Name = "John"
	u2.Age = 30
	u2.Email = "john.doe@example.com"

	fmt.Println(u1) // Still Anna
	fmt.Println(u2) // John Doe

	// Change name of origin by using pointers
	rename(&u1, "Derek")
	fmt.Println(u1)

	u3 := User{
		Name: "Nik",
		Age:  29,
		Address: Address{
			City:    "Batumi",
			Country: "Georgia",
		},
	}
	fmt.Println(u3)
}

func rename(u *User, name string) {
	u.Name = name
}
