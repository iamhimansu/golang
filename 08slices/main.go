package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	Slices
	*/

	slice := []string{"Mango", "Litchi"}

	fmt.Printf("Type of slice is %T\n", slice)

	sort.Strings(slice)
	fmt.Println(slice)

	highScores := make([]int, 6) // returns type is same as the data type argument

	fmt.Println(highScores)

	/**
	Removing an element from a slice at a specific index
	*/

	var courses = []string{"React", "Vue", "PHP", "Swift"} //remove php

	fmt.Println(courses)

	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)

}
