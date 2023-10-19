package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://www.lco.dev"

func main() {
	fmt.Println("hello")

	resp, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T", resp)

	defer resp.Body.Close() //Caller's responsibility Close the connection

	dataBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)

}
