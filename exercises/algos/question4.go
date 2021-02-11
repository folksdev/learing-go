// https://www.careercup.com/question?id=5147882054221824

package main

import "fmt"

func toMap(slc1 []string) map[string]int {
	keys := make(map[string]int)
	for i, v := range slc1 {
		if _, ok := keys[v]; !ok {
			keys[v] = i
		} else {
			keys[v] += i
		}
	}
	return keys
}

func twoSum(nums []int, target int) []int {
	keys := make(map[int]int)
	for i, v := range nums {
		val := target - v
		if idx, ok := keys[val]; ok {
			return []int{idx, i}
		}

		keys[v] = i
	}

	return nil
}

func solve(slc [][]int) []string {
	for _, list := range slc {

	}

	asMap := toMap(slc)

	if idx, ok := asMap[w]; ok {
		sIdx, exIdx := 0, idx-n

		if exIdx > sIdx {
			sIdx = exIdx
		}

		return slc[sIdx : idx+1]
	}
	return nil
}

func main() {
	var slc1 []string = []string{"aaa", "bbb", "ccc", "booking", "alpha", "beta", "gamma"}

	fmt.Println("Input: ", slc1)

	fmt.Println("k=3, w=booking: ", solve(slc1, "booking", 3))
	fmt.Println("k=2, w=beta: ", solve(slc1, "beta", 2))
	fmt.Println("k=10, w=gamma: ", solve(slc1, "gamma", 10))
	fmt.Println("k=3, w=aaa: ", solve(slc1, "aaa", 3))

}
