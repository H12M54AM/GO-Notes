package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Adds standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")

	//Reads the string or input, makes output 
	input, _ := reader.ReadString('\n')
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