package solution

import (
	"math"
)

func SecondSolution(nthPrimeNumber int) int {
	if nthPrimeNumber < 1 {
		return 0
	}

	var result int

	for number, primesCount := 2, 0; primesCount < nthPrimeNumber; number++ {
		if isPrimeNumberSecondSolution(number) {
			primesCount++
			result = number
		}
	}

	return result
}

func isPrimeNumberSecondSolution(number int) bool {
	// We can use more efficient algorithm in order to find a prime number through a square of a number
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	// 1 is not a prime number
	return number > 1
}
