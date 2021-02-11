package main

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func (e *Employee) PrintInfo() {
	fmt.Println("===================")
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println("===================")
}

func main() {
	var emp1 *Employee             // emp1: nil
	emp1 = new(Employee)           // emp1: &{'' ''  0}
	emp1.FirstName = "FolksDev"    // emp1: &{FolksDev ''  0}
	emp1.LastName = "FolksDevOglu" // emp1: &{FolksDev FolksDevOglu  0}
	emp1.Age = 1                   // emp1: &{FolksDev FolksDevOglu  1}

	emp1.PrintInfo() /*
		===================
		FolksDev FolksDevOglu
		1
		===================
	*/

	emp4 := Employee{"SomeOne", "", 22}
	emp5 := emp4

	emp4.FirstName = "AnotherOne"
	emp5.Age = 39
	emp4.PrintInfo()
	emp5.PrintInfo()

	fmt.Println("\n")
	emp6 := &Employee{"SomeOne", "", 22}
	emp7 := emp6

	emp6.FirstName = "AnotherOne"
	emp7.Age = 39
	emp6.PrintInfo()
	emp7.PrintInfo()

}
