package main

import (
	"fmt"
)

func main() {
	//number := 20
	//fmt.Println(reflect.TypeOf(number))
	//fmt.Println(reflect.TypeOf(&number))
	//fmt.Printf("Type of %T\n", &number)
	//
	//floatNumber := 20.0
	//fmt.Println(reflect.TypeOf(floatNumber))
	//fmt.Println(reflect.TypeOf(&floatNumber))
	//fmt.Printf("Type of %T\n", &floatNumber)
	//
	//stringVal := "reflect"
	//fmt.Println(reflect.TypeOf(stringVal))
	//fmt.Println(reflect.TypeOf(&stringVal))
	//fmt.Printf("Type of %T\n", &stringVal)
	//
	//boolVal := false
	//fmt.Println(reflect.TypeOf(boolVal))
	//fmt.Println(reflect.TypeOf(&boolVal))
	//fmt.Printf("Type of %T\n", &boolVal)

	//number := 20
	//numberPointer := &number
	//fmt.Println(number)
	//fmt.Println(numberPointer)
	//fmt.Println(*numberPointer)
	//
	//*numberPointer = 40
	//fmt.Println(number)
	//fmt.Println(*numberPointer)
	//
	//number = 90
	//fmt.Println(number)
	//fmt.Println(*numberPointer)

	myFloatPointer := returnPointer()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myFloatPointer = returnPointer()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myFloatPointer = returnPointer()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	for i := 0; i < 100; i++ {
		myFloatPointer = returnPointer()
		fmt.Println(myFloatPointer)
		fmt.Println(*myFloatPointer)
	}

}

func returnPointer() *float64 {
	var myFloat = 32.5
	return &myFloat
}
