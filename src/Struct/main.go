package main

import ("fmt")

func main() {
	fmt.Println(" ") // Spaces
	pet := Dog{
		"poodle", 
		23,
		"woof",
	}
	pet.Speak()
	spet := Cat{"Sphix", 34}
	fmt.Println("Primary Pet Info:", pet)

	fmt.Printf("%+v\n", pet)
	fmt.Println(" ") // Spaces

	fmt.Println("Secondary Pet Info:", spet)
	fmt.Printf("%+v\n", spet)
	
	fmt.Println(" ") // Spaces
}
// Custom
type Dog struct {
	Breed string
	Age int64
	Sound string
}
// Custom 
type Cat struct {
	Breed string
	Age int64
}

// d Dog is a reciever, d is used to access the Dog struct 
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}