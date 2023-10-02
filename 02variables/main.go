package main

import "fmt"

const (
	GithubToken = "gph_126845084bvt0aH4FDP6HL3157GDqacr" // this is available publicly
)

func main() {
	var username = "Himanshu Raj Aman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	/**
	Small integers uint8 -> 8bit 1byte 0-255
	Medium integers uint16 -> 16 bit 2bytes 0-65535
	Large integers uint32 -> 32 bit 4bytes 0-4294967295
	Large integers uint64 -> 64 bit 8 bytes 0-18446744073709551615
	*/
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var medVal uint16 = 65535
	fmt.Println(medVal)
	fmt.Printf("Variable is of type: %T \n", medVal)

	var largeVal uint32 = 4294967295
	fmt.Println(largeVal)
	fmt.Printf("Variable is of type: %T \n", largeVal)

	var extraLargeVal uint64 = 18446744073709551615
	fmt.Println(extraLargeVal)
	fmt.Printf("Variable is of type: %T \n", extraLargeVal)

	/**
	Default values of integers and strings
	*/
	var declaredInt int //default -> 0
	fmt.Println(declaredInt)
	fmt.Printf("Variable is of type: %T \n", declaredInt)

	var declaredString string //default -> ''
	fmt.Println(declaredString)
	fmt.Printf("Variable is of type: %T \n", declaredString)

	/**
	Walrus operators
	* This operator is only allowed inside methods
	*/
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(GithubToken)
	fmt.Printf("Variable is of type: %T \n", GithubToken)

}
