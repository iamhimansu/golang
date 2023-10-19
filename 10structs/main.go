package main

import "fmt"

func main() {

	fmt.Print("Structs here")
	// GO Lang does not provide support inheritance
	// no parent, or super is present here like Java
	// The go developers thought that through inheritance, code becomes complex

	user := User{
		Name:   "Himanshu",
		Email:  "test@mail.com",
		Status: false,
		Age:    21,
	}

	fmt.Println(user)
	/**
	To get all details, use %+v  in printf
	*/
	fmt.Printf("User details: %+v", user)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
