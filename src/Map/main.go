package main

import (
	"fmt"
	"sort"
	"time"
) 

func main() {
	fmt.Println(" ")

	//Like a dictionarie in Python, you set your keys and your values to whatever types you want
	prov := make(map[string]string)
	prov ["BC"] = "British Columbia"
	prov ["AB"] = "Alberta"
	prov ["SA"] = "Saskatewan"
	prov ["MA"] = "Manitoba"
	fmt.Println(prov)
	fmt.Println(" ")
	
	where := prov["BC"]
	fmt.Println(where)
	fmt.Println(" ")
	
	//If I want to list it
	for k, v := range prov{
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println(" ")
	
	//If I wanted in order, I must convert the map into a slice, then sort it
	keys := make([]string, len(prov))
	i := 0
	for k := range prov {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	fmt.Println(" ")
	
	//Time to sort
	sort.Strings(keys)
	fmt.Println(keys)
	fmt.Println(" ")
	for i := range keys {
		fmt.Println(prov[keys[i]])
	}
	fmt.Println(" ")
	t := time.Now()
	fmt.Println("Tested at", t.Format(time.Kitchen))
	fmt.Println(t.Format(time.RFC850))
	fmt.Println(" ")
}