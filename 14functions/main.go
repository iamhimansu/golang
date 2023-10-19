package main

import "fmt"

func main() {
	fmt.Println("Functions in go")

	fmt.Println(add(50, 909))

	fmt.Println(addn(50, 80, 90, 67, 56))
}

func add(a, b int) int {
	return a + b
}

/*
*
Variadic function
*/
func addn(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
