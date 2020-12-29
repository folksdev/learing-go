package main

import "fmt"

func main() {

	// continue skips to the head of the loop block without executing the rest
	for x := 0; x <= 10; x++ {
		fmt.Println("before continue")
		continue
		fmt.Println("after continue")
	}
	fmt.Println("after continue example")

	// break skips the entire loop
	for x := 0; x <= 10; x++ {
		fmt.Println("before break")
		break
		fmt.Println("after break")
	}
	fmt.Println("after break example")
}
