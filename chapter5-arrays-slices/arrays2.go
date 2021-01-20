package smt2

import "fmt"

func smt2a() {
	var myIntArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// myIntArray2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var myFloatArray [5]float64
	// var myBooleanArray [4]bool
	// var myStringArray [2]string

	//fmt.Println(myIntArray)
	//fmt.Printf("Your array is: %v\n", myIntArray2)
	// fmt.Println(myFloatArray)
	// fmt.Println(myBooleanArray)
	// fmt.Println(myStringArray)

	length := len(myIntArray)
	for i := 0; i < length; i++ {
		myIntArray[i] = 10
	}

	// fmt.Println(myIntArray)

	for _, value := range myIntArray {
		fmt.Printf("value: %d\n", value)
	}

}
