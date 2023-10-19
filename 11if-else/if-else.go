package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else")

	var loginCount int = 18

	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount < 20 {
		fmt.Println("++Regular User")
	} else {
		fmt.Println("Addictive User")
	}

	//inline variable declaration
	//
	if result := 90; result < 50 {
		fmt.Println("Result is less")
	} else {
		fmt.Println("Result is more")
	}
}
