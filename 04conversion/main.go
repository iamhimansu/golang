package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a rating for our pizza:")

	//Comma, error syntax
	rating, _ := reader.ReadString('\n')

	fmt.Println("Thank-you for rating us", rating)
	fmt.Printf("Type of this rating is %T\n", rating)
	/** Convert string to integer to add 1 */

	ratingInt, error := strconv.ParseInt(strings.TrimSpace(rating), 0, 64)

	if error != nil {
		println("Something went wrong while converting string to int64: ", error)
	} else {
		ratingInt++
		println("Added +1 to your rating", ratingInt)
	}
}
