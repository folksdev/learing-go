package main

import "fmt"

func main() {
	//fmt.Printf("Value of 1/3: %f\n", 1.0/3.0)

	//formatted := fmt.Sprintf("Value of 1/3: %f\n", 1.0/3.0)
	//fmt.Println("String", formatted)

	//fmt.Printf("Value of product %s is %f TL\n", "Apple", 5.0)

	//fmt.Printf("Value of product %s is %f TL. Stock amount: %d\n", "Apple", 5.0, 250)

	// %s -> string
	// fmt.Printf("Value of product %s\n", "Apple")
	//
	// // %f -> float
	// fmt.Printf("Value of product %f\n", 250.0)
	//
	// // %d -> int
	// fmt.Printf("Value of product %d\n", 213)
	//
	// // %t -> boolean
	// fmt.Printf("Value of bool true %t\n", true)
	// fmt.Printf("Value of bool false %t\n", false)

	// %v -> any value
	//fmt.Printf("Values: %v, %v, %v, %v\n", 15, "Elma", 0.5*5.0, true)
	//fmt.Printf("Values: %v, %v, %v, %v\n", 15, "\n", 0.5*5.0, true)

	// %#v -> any value
	// fmt.Printf("Values: %#v, %#v, %#v, %#v\n", 15, "Elma", 0.5*5.0, true)
	// fmt.Printf("Values: %#v, %#v, %#v, %#v\n", 15, "\n", 0.5*5.0, true)

	// %T -> type of value
	//fmt.Printf("Values: %T, %T, %T, %T\n", 15, "\n", 0.5*5.0, true)

	// %
	//fmt.Printf("Values: %%\n")

	//fmt.Printf("%12s | %s\n", "Product", "Price")
	//fmt.Println("---------------------------------")
	//fmt.Printf("%12s | %2d\n", "Apple", 50)
	//fmt.Printf("%12s | %2d\n", "Folksdev", 5)
	//fmt.Printf("%12s | %2d\n", "Orange", 99)
	//fmt.Printf("%12s | %2d\n", "OrangeOrangeAB", 1299)

	fmt.Printf("Value of 1/3: %.3f\n", 1.0/3.0)
	fmt.Printf("%%1.2f: %1.2f\n", 12.1234567)
}
