package main // package block

import "fmt"

var name = "folksdev" // declare global variable

func main() { // function block

	var name2 = "asdf "
	fmt.Println(name)
	fmt.Println(name2)

	// declare function variable
	aVal := 5
	fmt.Println(aVal)
	if true { // if block

		name2 := 20
		fmt.Println(name)
		fmt.Println(name2)

		conditionVal := 20 // in codition scope
		fmt.Println(conditionVal)

		aVal := "some str"
		fmt.Println(aVal)

		// print global var
		// print function var

	}

	// print global var
	// print function var
	// print condition var
}
