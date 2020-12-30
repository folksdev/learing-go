// This file won't compile
package main

import "fmt"

func main() {
	aVal := "Pass!"
	fmt.Println(aVal)
	bVal, aVal := 2, "Fail"

	fmt.Println(bVal)
	fmt.Println(aVal)

	cVal, bVal, aVal := 1, 2, 3

	fmt.Println(cVal)
	fmt.Println(bVal)
	fmt.Println(aVal)
}
