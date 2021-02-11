/*
Write a function

`wordcount(str string) map[string]int`

which will return the number of occurences for each word in given string
String must be cleaned from punctuation marks.
For ex:

Input:  "hello. hello! world? folksdev, world21"
Output  [folksdev:1 hello:2 world:1 world21:1]

*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "hello. hello! world? folksdev, world21"

	fmt.Println("Input: ", str)

	fmt.Println("Output: ", wordcount(str))

}

func wordcount(str string) map[string]int {
	reg, err := regexp.Compile("[^\\w\\s]")

	if err != nil {
		return nil
	}

	str = reg.ReplaceAllString(str, "")
	splitted := strings.Split(str, " ")

	keys := make(map[string]int)
	for _, v := range splitted {
		if _, ok := keys[v]; ok {
			keys[v]++
		} else {
			keys[v] = 1
		}
	}

	return keys
}
