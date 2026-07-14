package solution

func FirstSolution(exceedNumber int) int {
	var result int

	first := 1
	second := 2

	for second < exceedNumber {
		if second%2 == 0 {
			result += second
		}

		next := first + second
		first = second
		second = next
	}

	return result
}
