package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	width := 4.75
	height := 3.0
	paintNeeded, err := calculatePaint(width, height)
	fmt.Println("Paint needed: ", paintNeeded, "lt", "Err:", err)

	width = 5.5
	height = 3.5
	paintNeeded, err = calculatePaint(width, height)
	fmt.Println("Paint needed: ", paintNeeded, "lt", "Err:", err)

	width = 5.5
	height = -3.5
	paintNeeded, err = calculatePaint(width, height)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Paint needed: ", paintNeeded, "lt", "Err:", err)
}

func calculatePaint(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		err := errors.New("Arguments cannot be negative")
		return 0, err
	}
	paintNeeded := (width * height) / 10
	return paintNeeded, nil
}
