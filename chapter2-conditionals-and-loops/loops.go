package main

import "fmt"

func main() {
	for x := 4; x <= 6; x++ {
		fmt.Println("X is now: ", x)
		y := x + 10
		fmt.Println("Y is now: ", y)
	}

	// Is x available here?
	x := "sadsadass"
	fmt.Println("X is now: ", x)

	// countdown  loop
	for x := 10; x >= 0; x-- {
		fmt.Println("X is now: ", x)
	}

	// Infinite loop example
	/*for y := 10; 1 == 1; y-- {
		fmt.Println("y is now: ", y)
	}
	fmt.Println("X is now: ", x)*/

	// Unreachable loop example
	for y := 10; false; y-- {
		fmt.Println("y is now: ", y)
	}
}
