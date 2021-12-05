package main

/* Make a script that asks the user what they want to type, 
and put that into an array. That way, thier text will be 
stored for later use, hopefully */

//Copy of Input.go

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	Text()
}

func Text() {
	//Adds standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")

	//Reads the string or input
	input, _ := reader.ReadString('\n')
	//output
	fmt.Println("You Entered:", input)

	fmt.Print("Enter Num: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err :=  strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		 fmt.Println("Value of Number: ", aFloat)
	}
}