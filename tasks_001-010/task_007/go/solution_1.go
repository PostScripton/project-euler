package solution

func FirstSolution(nthPrimeNumber int) int {
	if nthPrimeNumber < 1 {
		return 0
	}

	var result int

	for number, primesCount := 2, 0; primesCount < nthPrimeNumber; number++ {
		if isPrimeNumberFirstSolution(number) {
			primesCount++
			result = number
		}
	}

	return result
}

func isPrimeNumberFirstSolution(number int) bool {
	// Prime number is a number that can only be divided by either 1 or itself
	// thus, we exclude those numbers from loop range
	for i := 2; i <= number-1; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
