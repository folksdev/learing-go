package main

import (
	"fmt"
)

func main() {

	number := 20
	//fmt.Println(number)
	//passByValue(number)

	fmt.Println(number)
	passByRefence(&number)
	fmt.Println(number)

}

func passByValue(number int) int {
	number = number * 2
	fmt.Println(number)
	return number
}

func passByRefence(numberPoint *int) {
	*numberPoint *= 2
	fmt.Println(*numberPoint)
}
