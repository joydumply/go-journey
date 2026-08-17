package main

import "fmt"

func main() {

	fmt.Println("for loop")
	// Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}
	fmt.Println("========\n")

	// While loop
	fmt.Println("while loop")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("========\n")

	// Break and Continue
	fmt.Println("break & continue")

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("========")

	// Outer
	fmt.Println("outer label")

outer:
	for i := 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if j == 1 {
				continue outer // stops j loop and go to i loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("========")

	// for range collection iteration
	fmt.Println("for range collection iteration ")
	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Println(i, v) // key value
		nums[i] *= 10     // will change nums values, not a copy
	}

	fmt.Println(nums)

	fmt.Println("========")

	// Maps
	fmt.Println("Maps")
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}

}
