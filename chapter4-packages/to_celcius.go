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

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) // use buffered reader to get the data from the user
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return float, nil
}

func main() {

	fmt.Println("Enter Fahreheit: ") // Ask user to enter input
	fahrenheit, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("Fahrenheit: %.2f -> Celcius: %.2f\n", fahrenheit, celcius)
}
