package average

func Average(numbers ...float64) float64 {
	count := 0
	var total float64 = 0
	for _, number := range numbers {
		count++
		total += number
	}
	if count == 0 {
		return 0
	} else {
		return total / float64(count)
	}
}
