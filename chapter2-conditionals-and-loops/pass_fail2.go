package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sayı girin")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	var message string
	if grade >= 60 {
		message = "Geçti"
	} else {
		message = "Kaldı"
	}

	fmt.Println(message)

}
