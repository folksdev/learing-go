/*
	User enters a grade. If the grade is below 60 then `fail` will be printed
	otherwise `pass` will be printed
*/

package main

import (
	"bufio" // import to read input from the user
	"fmt"
	"log"
	"os" // import to read input from the user
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter grade: ")        // Ask user to enter input
	reader := bufio.NewReader(os.Stdin) // use buffered reader to get the data from the user
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
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
