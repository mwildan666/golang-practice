package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println("Default Months =", months)
	var slice1 = months[4:7]
	fmt.Println("Sliced Months =", slice1)
	fmt.Println("Slice Length =", len(slice1))
	fmt.Println("Slice Capacity =", cap(slice1))
	slice1[0] = "Deadcember"
	fmt.Println(months, "After Changing Slice Data")

	slice2 := append(slice1, "Diecember", "Doomcember")
	fmt.Println(slice2)

	slice3 := make([]string, len(slice2), cap(slice2))
	fmt.Println(slice3, "Before Being Copied")
	copy(slice3, slice2)
	fmt.Println(slice3, "After Being Copied")

}
