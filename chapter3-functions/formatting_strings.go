package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "Ürünler", "Fiyat")
	fmt.Println("---------------------------------")
	fmt.Printf("%12s | %2d\n", "Elma", 50)
	fmt.Printf("%12s | %2d\n", "Armut", 5)
	fmt.Printf("%12s | %2d\n", "Portakal", 99)

	fmt.Printf("%%3.2f: %3.2f\n", 99.0)
	fmt.Printf("%%7.3f: %7.3f\n", 12.1234567)

}
