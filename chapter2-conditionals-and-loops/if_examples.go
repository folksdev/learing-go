package main

import (
	"fmt"
)

func main() {

	if 1 < 2 {
		fmt.Println("1 < 2 is true")
	}

	if 3 < 2 {
		fmt.Println("3 < 2 is true")
	}

	// run if condition is true
	if true {
		fmt.Println("I will be printed")
	}

	// not run if condition is false
	if false {
		fmt.Println("I won't be printed")
	}

	// multiple branches with else if
	grade := 35

	if grade == 100 {
		fmt.Println("Perfect")
	} else if grade > 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// negating boolean
	if !true {
		fmt.Println("I won't be printed")
	}

	if !false {
		fmt.Println("I will be printed")
	}

	// logical operators

	if true && true {
		fmt.Println("I will be printed")
	}

	if true && false {

	}

	if false && true {

	}

	if true && !false {

	}

	if !false && !false {

	}

	if !true && !false {

	}

	if true || true {
		fmt.Println("I will be printed")
	}

	if true || false {

	}

	if false || true {

	}

	if true || !false {

	}

	if !false || !false {

	}

	if !true || !false {

	}

}
