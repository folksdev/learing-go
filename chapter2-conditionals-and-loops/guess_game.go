package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	rounds := 3

	target := rand.Intn(100) + 1
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)

	isCorrect := false

	for x := 0; x < rounds; x++ {
		userRound := rounds - x
		fmt.Println("Sayıyı tahmin etmeye çalışın. Mevcut Hakkınız ", userRound)

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			isCorrect = true
			break
		} else if guess > target {
			if userRound != 1 {
				fmt.Println("Aşağı")
			}
		} else {
			if userRound != 1 {
				fmt.Println("Yukarı")
			}
		}
	}

	if isCorrect {
		fmt.Println("Tebrikler doğru bildin")
	} else {
		fmt.Println("Bilemedin")
	}

}
