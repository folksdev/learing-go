package shared

import (
	"bufio" // import to read input from the user
	"os"    // import to read input from the user
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) // use buffered reader to get the data from the user
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return float, nil
}
