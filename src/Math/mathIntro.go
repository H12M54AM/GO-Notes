package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	check()

	/* i1, i2, i3 := 21, 23, 89
	intSum := i1 + i2 + i3
	fmt.Println("The sum of the Integer is:", intSum)

	f1, f2, f3 := 23.2, 34.2, 98.4
	floatSum := f1 + f2 + f3
	fmt.Println("The sum of the Flaot is:", floatSum)
	
	//Because floatSum was already delcared, it can be used again without the use of the colon
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The Rounded Float is:", floatSum) */

	circleRad := 15.5
	circumference := circleRad * 2 * math.Pi
	fmt.Printf("The Circumference is: %.2f\n", circumference) 
	/* What the "%.2f" mean is the floating
	number will have 2 digit decimals affter the number */

	/* Important when doing that is to keep the "%.2f" in the strings area as if you were typring 
	something out */
}

//Asks what the user would like to input as a float
func check() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Place your number:")
	
	input, _ := reader.ReadString('\n')
	fmt.Print("That's", input, ", correct?")

	if input == "yes" {
		fmt.Println("Your number is", input)
	} else {
		fmt.Print("check what you typed, please")
	}

	numInput, _ := reader.ReadString('\n')
	aFloat, err :=  strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	fmt.Println(aFloat)
	
	if err != nil {
		fmt.Println(err)
	} else {
		 fmt.Println("Value of Number: ", aFloat)
	}
}

// this file needs work..