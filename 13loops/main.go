package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")

	//days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	/**
	Remember in arrays, if any index is blank, length will always show the total length of the array
	Thus, it can break in for loop
	Always use []Slices{}
	*/
	/*for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}*/

	//Another method with index
	/*for index := range days {
		fmt.Println(days[index])
	}*/

	//Another method with index, value
	/*for _, day := range days {
		fmt.Println(day)
	}*/

	//There is no while loop in golang
	//We can use for loop to simulate a while loop
	//Also there is no preincrement in golang
	//Example:
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	/*limit := 0
	for limit < 10 {
		fmt.Println(limit)
		limit++
	}*/

	//Breaks
	limit := 10
	for limit < 10 {
		if limit == 2 {
			break
		}

		if limit == 5 {
			limit++
			continue
		}
		if limit == 6 {
			goto exit
		}
		limit++
	}
exit:
	fmt.Println("Exiting...")
}
