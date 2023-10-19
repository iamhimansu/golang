package main

import (
	"fmt"
	rand2 "math/rand"
)

func main() {
	fmt.Println("Switch cases")

	diceNumber := rand2.Intn(6) + 1 // 6 is exclusive as of half interval
	fmt.Printf("You got %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("You can open up and move %v places", diceNumber)
	case 2:
		fmt.Printf("You can move %v places", diceNumber)
	case 3:
		fmt.Printf("You can move %v places", diceNumber)
	case 4:
		fmt.Printf("You can move %v places\n", diceNumber)
		//Fallthrough: this will go to the next condition and return the result,
		//the condition will always be true
		fallthrough
	case 5:
		fmt.Printf("You can move %v places", diceNumber)
	case 6:
		fmt.Printf("You can move %v places, and again roll the dice 🙂", diceNumber)
	}
}
