package main

import (
	"average"
	"fmt"
	"txtreader"
)

func main() {
	var meatConsumptionArray []float64 = txtreader.ReadTxtFile("data.txt")

	avg := average.CalculateAverage(meatConsumptionArray)

	fmt.Printf("Your weekly average is %.2f kg\n", avg)
}
