package solution

import (
	"math"
)

func ThirdSolution(nthPrimeNumber int) int {
	if nthPrimeNumber < 1 {
		return 0
	}

	result := 2

	// As long as all prime numbers are odd except 2,
	// we can iterate over only odd numbers and count 2 as already counted value
	for number, primesCount := 3, 1; primesCount < nthPrimeNumber; number += 2 {
		if isPrimeNumberThirdSolution(number) {
			primesCount++
			result = number
		}
	}

	return result
}

func isPrimeNumberThirdSolution(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return number > 1
}
