package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	/*var myIntNumber int = 5
	var myFloatNumber float64 = 4.6

	fmt.Println(float64(myIntNumber) + myFloatNumber)*/

	//Random numbers
	//var random int = rand.Intn(5) //Open half interval, range is always exclusive
	//fmt.Println(random)

	//Crypto random number
	random, _ := rand.Int(rand.Reader, big.NewInt(590678543243))

	fmt.Println(random)

}
