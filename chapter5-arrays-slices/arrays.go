package smt

import "fmt"

func smt() {
	var myIntArray [10]int
	// var myFloatArray [5]float64
	// var myBooleanArray [4]bool
	// var myStringArray [2]string

	// fmt.Println(myIntArray)
	// fmt.Println(myFloatArray)
	// fmt.Println(myBooleanArray)
	// fmt.Println(myStringArray)

	firstInt := myIntArray[0]
	fmt.Println(firstInt)
	myIntArray[0] = 1
	fmt.Println(myIntArray[0])
	fmt.Println(myIntArray)

	myIntArray[2] = 10
	myIntArray[3] = 10

	for i := 0; i < 10; i++ {
		myIntArray[i] = 10
	}
	fmt.Println(myIntArray)
}
