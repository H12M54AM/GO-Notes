package main

import ("fmt")

func main() {
	val := 32
	p := &val // &val directs to the memory address not the actual value

	fmt.Println(" The val is:", *p)
}