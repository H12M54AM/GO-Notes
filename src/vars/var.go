package main

import ("fmt")

var name string = "Edward"

//Checks if name is correct
func main() {
	if (name == "Edward") {
		fmt.Println(name)
	} else {
		fmt.Println("there is no name")
	}
}