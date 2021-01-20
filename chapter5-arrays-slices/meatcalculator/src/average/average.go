package average

func CalculateAverage(array []float64) float64 {
	sum := 0.0

	for _, value := range array {
		sum += value
	}
	average := sum / float64(len(array))
	return average
}
