/*
Write a function

`intersection(slc1 []int, slc2 []int) []int`

which will return the intersection of given slices

For ex:

Input:  {1, 2, 3, 4}, {2, 3, 5, 8}
Outout  {2, 3}

*/

package main

import "fmt"

func main() {
	slc1, slc2 := []int{1, 2, 3, 4}, []int{2, 3, 5, 8}

	fmt.Println("Slice1: ", slc1, "Slice2: ", slc2)

	fmt.Println("Output: ", intersection(slc1, slc2))

}

func toMap(slc []int) map[int]int {
	keys := make(map[int]int)
	for i, v := range slc {
		if _, ok := keys[v]; !ok {
			keys[v] = i
		}
	}
	return keys
}

func intersection(slc1 []int, slc2 []int) []int {

	slcMap := toMap(slc2)
	copy := make([]int, 0, 20)
	for _, v := range slc1 {
		if _, ok := slcMap[v]; ok {
			copy = append(copy, v)
		}
	}

	return copy
}
