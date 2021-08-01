package main

import "fmt"

func main(){
	// Basic Constant
	const firstName string = "Muhammad Wildan"
	const lastName string = "Anandasyah"
	const age = 23

	// Will Not Change
	// firstName = "Wil Wil"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// Multiple Constant
	const (
		herFName = "Adelia Utami"
		herLName = "Nasution"
		herAge = 23
	)

	fmt.Println(herFName)
	fmt.Println(herLName)
	fmt.Println(herAge)
}