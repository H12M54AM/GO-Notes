package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hello from go"
	file, err := os.Create("./Thefirstfile.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./Thefirstfile.txt")
}

// Used multiple times 
func checkError(err error) {
	if err != nil {
		panic("Something wrong happend here")
	}
}

// Will read the name of said file
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text Read from file", string(data))
}