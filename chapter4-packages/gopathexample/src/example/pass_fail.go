/*
	User enters a grade. If the grade is below 60 then `fail` will be printed
	otherwise `pass` will be printed
*/

package main

import (
	// import to read input from the user
	"fmt"
	"log" // import to read input from the user

	"github.com/folksdev/keyboard"
)

func main() {

	fmt.Println("Enter grade: ") // Ask user to enter input
	grade, err := keyboard.ReadFloat()

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "Pass!"
	} else {
		status = "Fail!"
	}
	fmt.Println(status)

}
