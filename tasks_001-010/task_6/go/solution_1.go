package solution

func FirstSolution(until int) int {
	var (
		sumOfSquares int
		squareOfSum  int
	)

	for i := 1; i <= until; i++ {
		sumOfSquares += i * i
		squareOfSum += i
	}

	squareOfSum *= squareOfSum

	return squareOfSum - sumOfSquares
}
