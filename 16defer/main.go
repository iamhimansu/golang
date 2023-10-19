package main

import "fmt"

func main() {

	fmt.Println("Defer in go")

	/**
	In go, we can control the execution of the expression or methods
	Using "defer" keyword

	The line having defer keyword will be executed at the very end of the program
	Example:
	*/

	defer fmt.Println("Hello")
	fmt.Println("World")

	/**
	It follows LIFO
	*/

	defer fmt.Println("Executing...")
	defer print5Numbers()

}

func print5Numbers() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
