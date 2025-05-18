package solution

func FirstSolution(number int) int {
	var largest int

	for number > 1 {
		pf := primeFactor(number)
		number /= pf

		if pf > largest {
			largest = pf
		}
	}

	return largest
}

func primeFactor(number int) int {
	for divider := 2; divider <= number; divider++ {
		if number%divider == 0 {
			return divider
		}
	}

	return 1
}
