package main

import (
	"fmt"
)

func main() {
	var myIntArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice1 := myIntArray[:]
	mySlice2 := myIntArray[1:5]
	mySlice3 := myIntArray[:5]
	mySlice4 := myIntArray[3:]

	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)
	fmt.Println(mySlice4)

	mySlice5 := []string{"a", "b", "c"}
	fmt.Println(mySlice5)

	mySlice6 := make([]float64, 4, 4)
	fmt.Println(mySlice6)

	mySlice7 := make([]float64, 4)
	fmt.Println(mySlice7)

	mySlice8 := make([]float64, 2, 4)
	fmt.Println(mySlice8)

	fmt.Println(len(mySlice8))
	fmt.Println(cap(mySlice8))

	mySlice8 = append(mySlice8, 54.0)

	fmt.Println(mySlice8)

	fmt.Println(len(mySlice8))
	fmt.Println(cap(mySlice8))

	mySlice8 = append(mySlice8, 64.0)
	mySlice8 = append(mySlice8, 74.0)

	fmt.Println(mySlice8)

	fmt.Println(len(mySlice8))
	fmt.Println(cap(mySlice8))

	// func make([]T, len, cap) []T
}
