package main

import (
	"fmt"
	"time"
)

func main() {
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 
	for i := range value {
		fmt.Println(value[i])
	}
	fmt.Println(" ")

	num := 0
	// 200 is the max value set
	for num < 200{
		num++
		fmt.Println("Num:", num)
		if num == 150 {
			// A saftey precaution since it is a loop
			goto stop
		}
	}
	fmt.Println(" ")
	fmt.Println(time.Now().Format(time.ANSIC))
	
	stop : fmt.Println("Program stopped")
	
}