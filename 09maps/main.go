package main

import (
	"fmt"
)

func main() {

	/**
	Stores key and value pairs
	*/

	courses := make(map[string]string, 5)

	courses["JS"] = "Javascript"
	courses["PHP"] = "PHP"
	courses["SW"] = "Swift"
	courses["RB"] = "Ruby"
	courses["GO"] = "Golang"

	fmt.Println(courses)

	/**
	delete a key from map
	*/
	delete(courses, "JS")
	for _, value := range courses {
		fmt.Printf("Value is %v\n", value)
	}
}
