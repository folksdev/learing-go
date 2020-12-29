/*
	User enters grades, When entering data is finished, user has to send C (in capital)
	to start calculation. If the average is below 60 then `fail` will be printed
	otherwise `pass` will be printed
*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// Guess the return values

	bool, err := strconv.ParseBool("true")
	fmt.Println(bool, err)

	bool, err = strconv.ParseBool("false")
	fmt.Println(bool, err)

	bool, err = strconv.ParseBool("abs")
	fmt.Println(bool, err)

	file, ioErr := os.Open("pass_fail.go")
	fmt.Println(file.Name, ioErr)

	file, ioErr = os.Open("abs")
	fmt.Println(file, ioErr)

	response, httpErr := http.Get("http://play.golang.org")
	fmt.Println(response.Status, httpErr)

	response, httpErr = http.Get("http://invalidUrl")
	fmt.Println(response, httpErr)
}
