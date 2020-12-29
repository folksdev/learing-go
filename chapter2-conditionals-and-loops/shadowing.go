// This file won't compile
package main

import (
	"fmt"
)

func main() {
	var int int = 10 // Name int shadows type int
	fmt.Println(int)
	var count int = 20 // error: int is not a type
	fmt.Println(count)
}
