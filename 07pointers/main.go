package main

import "fmt"

func main() {

	var nullPtr *int //null pointer
	fmt.Println(nullPtr)

	var number int = 7

	var ptr = &number

	fmt.Println("Variable address is", ptr)

	*ptr = *ptr + 4
	fmt.Println("Value is", *ptr)
}
