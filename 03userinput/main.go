package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a rating for our pizza:")

	//Comma, error syntax
	rating, _ := reader.ReadString('\n')

	fmt.Println("Thankyou for rating us", rating)
	fmt.Printf("Type of this rating is %T", rating)

}
