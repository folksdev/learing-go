package txtreader

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadTxtFile(path string) []float64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	meatConsumptionSlice := []float64{}

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		meatConsumptionSlice = append(meatConsumptionSlice, value)
	}

	return meatConsumptionSlice
}
