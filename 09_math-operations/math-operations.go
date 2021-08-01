package main

import "fmt"

func main(){
	// Basic Operations
	var hasil = 10 + 10
	fmt.Println(hasil)

	var a = 38
	var b = 13
	var c = a - b
	fmt.Println(c)

	// Unary Operations
	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var pos = +100 // + is optional, by default number is positive
	var neg = -100
	fmt.Println(pos)
	fmt.Println(neg)
}