package shared

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
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
