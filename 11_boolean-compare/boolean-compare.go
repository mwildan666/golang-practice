package main

import "fmt"

func main() {
	var val1 = 75
	var val2 = 10

	var qualif1 = val1 >= 80
	var qualif2 = val2 >= 10

	var qualified bool = qualif1 && qualif2
	var cheat bool = qualif1 || qualif2
	var bribe bool = !qualified

	fmt.Println(qualified)
	fmt.Println(cheat)
	fmt.Println(bribe)
}
