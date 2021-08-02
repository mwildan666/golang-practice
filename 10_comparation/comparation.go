package main

import "fmt"

func main() {
	// Comparation Operator (String)
	var name1 = "Wildan"
	var name2 = "Wildan"
	fmt.Println(1, name1 == name2)

	// Comparation Operator (Number)
	var val1 = 100
	var val2 = 200
	fmt.Println(2, val1 == val2)
	fmt.Println(3, val1 < val2)
	fmt.Println(4, val1 > val2)
	fmt.Println(5, val1 <= val2)
	fmt.Println(6, val1 >= val2)
	fmt.Println(7, val1 != val2)
}
