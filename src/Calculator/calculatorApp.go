package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

//psuedocode/function of code
/*
	1. Once app starts, it will ask for an input
	2. if that input is a string, then panic
		There will be no error if the float is entered
	3. App will ask user for a second input
	4. If the input is a string, then panic
		There will be no error if the float is entered
	5. After the inputs were entered, then they will be added
		other operators are soon to come..
	6. Result will be displayed in terminal
	7. Timestamp will be displayed

*/
func main (){
	//Step 1.
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter First Value: ")

	convert, _:= reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(convert), 64)
	if err !=  nil{
		panic("Incorrect input, please try agian")
	}

	fmt.Print("Enter Second Value: ")
	convert2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(convert2), 64)
	if err != nil{
		panic("Incorrect Input, please try again")
	}

	sum := float1 + float2
	sum = math.Round(sum * 100) /100

	fmt.Println("The sum of", float1, "+", float2, "=", sum)
	begin := time.Now()
	fmt.Println("Posted on", begin.Format(time.RFC1123))
}