package solution

func FirstSolution(sum int) int {
	for a := 1; a <= sum; a++ {
		for b := a + 1; b <= sum; b++ {
			c := sum - a - b

			if (a*a)+(b*b) == (c * c) {
				return a * b * c
			}
		}
	}

	return 0
}
