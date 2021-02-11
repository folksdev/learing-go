/*
Write a function

`distinct(slc []int) []int`

which will return the distinct items in given slice

For ex:

Input:  {1, 1, 2, 2, 3, 3, 4, 4}
Output  {1, 2, 3, 4}

*/

package main

import "fmt"

func main() {
	var slc []int = []int{1, 1, 2, 2, 3, 3, 4, 4}

	fmt.Println("Input: ", slc)

	fmt.Println("Output: ", distinct(slc))

}

func distinct(slc []int) []int {
	copy := make([]int, 0, len(slc))
	keys := make(map[int]int)
	for i, v := range slc {
		if _, ok := keys[v]; !ok {
			keys[v] = i
		} else {
			copy = append(copy, v)
		}
	}

	return copy
}
