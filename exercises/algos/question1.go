// https://www.careercup.com/question?id=5106869203369984

// Full ref: https://www.careercup.com/page?pid=bookingcom-interview-questions

package main

import "fmt"

func toMap(slc1 []int) map[int]int {
	keys := make(map[int]int)
	for i, v := range slc1 {
		if _, ok := keys[v]; !ok {
			keys[v] = i
		}
	}
	return keys
}

func distinct(slc1 []int) []int {
	asMap := toMap(slc1)
	result := make([]int, 0, len(asMap))
	for k := range asMap {
		result = append(result, k)
	}

	return result
}

func intersection(slc1 []int, slc2 []int) []int {

	var long []int
	var short []int
	if len(slc1) < len(slc2) {
		long, short = slc2, slc1
	} else {
		long, short = slc1, slc2
	}

	longMap, shortMap := toMap(long), toMap(short)

	var result []int

	for s1 := range longMap {
		if _, ok := shortMap[s1]; ok {
			result = append(result, s1)
		}
	}
	return result
}

func main() {
	var slc1 []int = []int{2, 5, 3, 2, 8, 1}
	var slc2 []int = []int{7, 9, 5, 2, 4, 10, 10}
	var slc3 []int = []int{6, 7, 5, 5, 3, 7}

	var int1 = intersection(slc1, slc2)
	var int2 = intersection(slc2, slc3)
	var int3 = intersection(slc1, slc3)

	f := append(int1, int2...)
	f = append(f, int3...)
	f = distinct(f)

	fmt.Println("Input Arrays")
	fmt.Println(slc1)
	fmt.Println(slc2)
	fmt.Println(slc3)

	fmt.Println("Elements Present in At Least Two Arrays")
	fmt.Println(f)
}
