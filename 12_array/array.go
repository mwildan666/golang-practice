package main

import "fmt"

func main() {
	// String Array
	var name [3]string
	name[0] = "Muhammad"
	name[1] = "Wildan"
	name[2] = "Anandasyah"

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Integer Array
	var number = [5]int{
		10,
		15,
		20,
		25,
		30,
	}
	fmt.Println(number)
	fmt.Println(len(number))
}
