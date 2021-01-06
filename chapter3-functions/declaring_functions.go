package main

import "fmt"

func main() {
	repeatLine("Samet", 10)
}

func sayHi() {
	fmt.Println("Hi")
}

func sayHiName(name string) {
	fmt.Println("Hi", name)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
