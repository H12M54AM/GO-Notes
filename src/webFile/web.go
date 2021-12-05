package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://apod.nasa.gov/apod/astropix.html"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic("Something went wrong in Getting the information")
	}
	fmt.Printf("Response type: %T\n", resp)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error in Reading Bytes")
	}

	content := string(bytes)
	fmt.Print(content)
}
