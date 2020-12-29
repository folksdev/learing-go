package main

import (
	"fmt"
	"reflect" // import for checking types
	"time"    // import time package to use time.Time type
)

func main() {
	var now time.Time = time.Now()         // time.Now returns a time.Type value representing current date time
	var day, month, year = now.Date()      // Date returs a triplet for (day: Int, month: Month, year:Int)
	fmt.Println("Date:", day, month, year) // Print date with components
	fmt.Println("Type of day:", reflect.TypeOf(day))
	fmt.Println("Type of month:", reflect.TypeOf(month))
	fmt.Println("Type of year:", reflect.TypeOf(year))

	var monthAsInt int = int(month)           // get integer value of month
	var monthAsString string = month.String() // call string method of time.Month
	fmt.Println("Month ", monthAsString, " is the ", monthAsInt, "th month.")

	var january1 = time.January
	var january2 = time.Month(1) // not zero but starts from one
	fmt.Println("January1: ", january1, "January2: ", january2)
	fmt.Println("Month ", january1, " is the ", int(january1), "th month.")
	fmt.Println("Month ", january2, " is the ", int(january2), "th month.")

	var october = time.Month(10)
	fmt.Println("Month ", october, " is the ", int(october), "th month.")

	newNow := time.Now() // Now holds a time value
	fmt.Println("Time is ", newNow.Hour(), newNow.Minute(), newNow.Second())
}
