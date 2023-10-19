package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Working with files")

	file, err := os.Create("./myname.txt") //returns a pointer to the file

	if err != nil {
		panic(err) //stops the program
	}
	bytesWritten, err := file.WriteString("Hello, my name is Himanshu Raj Aman") //returns the size in bytes
	if err != nil {
		file.Close()
		panic(err)
	}

	file.Close()
	fmt.Println(bytesWritten)

	//Read file
	read, err := os.ReadFile("./myname.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(read))
}
