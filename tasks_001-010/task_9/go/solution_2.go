package solution

func SecondSolution(sum int) int {
	for a := 3; a <= (sum-3)/3; a++ {
		for b := a + 1; b <= (sum-a)/2; b++ {
			c := sum - a - b

			if (a*a)+(b*b) == (c * c) {
				return a * b * c
			}
		}
	}

	return 0
}
