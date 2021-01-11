package main

import (
	"fmt"
	"log"

	"./shared"
)

func main() {
	fmt.Println("Enter a grade")

	grade, err := shared.ReadFloat()

	if err != nil {
		log.Fatal(err)
	}

	var message string
	if grade >= 60 {
		message = "Pass"
	} else {
		message = "Fail"
	}

	fmt.Printf("Grade: %.2f Pass/Fail: %s\n", grade, message)

}
