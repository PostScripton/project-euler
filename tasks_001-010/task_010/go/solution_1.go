package solution

import (
	"math"
)

func FirstSolution(until int) int {
	var result int

	for number := 0; number <= until; number++ {
		if isPrimeNumber(number) {
			result += number
		}
	}

	return result
}

func isPrimeNumber(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return number > 1
}
