package solution

func FirstSolution(until int) int {
	var result int

	for num := 0; num < until; num++ {
		if num%3 == 0 || num%5 == 0 {
			result += num
		}
	}

	return result
}
