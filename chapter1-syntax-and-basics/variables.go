package main

import (
	"fmt"
)

func main() {

	// Explicit decleration
	var number int
	var decimal float64
	var str string

	// Assigning values
	number = 1
	decimal = 3.75
	str = "Hello form folksdev"

	// Print values stored in variables
	fmt.Println(number)
	fmt.Println(decimal)
	fmt.Println(str)

	// Assing value while defining
	var number2 int = 255
	var decimal2 float64 = 37.5
	var str2 = "Hello form folksdev 2"

	fmt.Println(number2)
	fmt.Println(decimal2)
	fmt.Println(str2)

	// Implicit decleration and assignment with :=
	number3 := 5
	decimal3 := 0.375
	str3 := "Hello form folksdev 3"

	fmt.Println(number3)
	fmt.Println(decimal3)
	fmt.Println(str3)
}
