package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" ")
	var colours [4]string 
	colours[0] = "red"
	colours[1] = "orange"
	colours[2] = "yellow"
	colours[3] = "green"
	fmt.Println("There are", colours, "and", len(colours), "colours in total")
	
	var cars = [2]string {"Mazda", "Honda"}
	fmt.Println(len(cars), "cars are within your garage")

	fmt.Println(" ")
	fmt.Println(" ")
	//Slice
	
	var companies = []string {"Apple", "Microsoft", "Red Bull", "Tesla", "Rimac"}
	fmt.Println("Companies: ", companies)
	fmt.Println(" ")
	
	//To add something
	companies = append(companies, "Dell")
	fmt.Println("Just to Append:", companies)
	fmt.Println(" ")
	
	//With a range
	companies = append(companies[1:len(companies)]) 
	//the range, if inputs are greater than one, can delete an item
	fmt.Println("With a range:", companies)
	fmt.Println(" ")

	//With a range that deletes the last item in array
	companies = append(companies[:len(companies)-1]) 
	
	fmt.Println("With a range that deletes the last item:", companies)
	fmt.Println(" ")
	fmt.Println(" ")
	
	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(" ")
}