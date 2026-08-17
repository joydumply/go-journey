package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("деление на ноль")
		return // this return will RETURN  err without a result (0)
	}
	result = a / b
	return // returns result without err
}

// variadic
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	res := add(2, 3) // 5
	fmt.Println(res)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("result: ", result)
	}

	s1 := sum(1, 2, 3)       // 6
	s2 := sum(1, 2, 3, 4, 5) // 15
	s3 := sum()              // 0

	fmt.Println(s1, s2, s3)
}
