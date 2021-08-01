package main

import "fmt"

func main(){
	// Basic Variable
	var name string

	name = "Muhammad Wildan Anandasyah"
	fmt.Println(name)

	// Automatic Variable
	var herName = "Adelia Utami Nasution"
	fmt.Println(herName)

	var ourAge int8 = 23
	fmt.Println(ourAge)

	nat := "Indonesia"
	fmt.Println(nat)

	// Multiple Variable
	var (
		day = 25
		month = "October"
		year = 2013
	)

	fmt.Println(day)
	fmt.Println(month)
	fmt.Println(year)
}