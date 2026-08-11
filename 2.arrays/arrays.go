package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Iterating array by for-loop
	for i := 0; i < len(primes); i++ {
		fmt.Println(primes[i])
	}

	fmt.Println("==========")

	// Range is better
	for _, value := range primes {
		fmt.Println(value)
	}

	fmt.Println("==========")

	// Checking array length
	fmt.Println("Array length: ", len(primes))
}
