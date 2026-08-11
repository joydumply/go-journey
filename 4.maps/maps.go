package main

import "fmt"

func main() {

	// Initialize map with make()
	m := make(map[string]int)

	// Set some data in map using key-value approach
	m["route"] = 66
	m["corner"] = 1
	fmt.Println(m)

	i := m["route"]
	j := m["corner"]
	k := m["hornet"]

	// Getting length of map
	n := len(m)

	// Removed item from map
	delete(m, "route")
	fmt.Println(i, j, k, n)
	fmt.Println(m)

	// Using Comma-Ok Idiom
	i, ok := m["route"]

	if ok {
		fmt.Println(i)
	} else { // It will be else
		fmt.Println("Route is not found")
	}

	m["route"] = 666

	// Iteration over the contents of a map
	for k, v := range m {
		fmt.Println("Key: ", k, "Value: ", v)
	}
}
