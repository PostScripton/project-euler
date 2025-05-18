package solution

import (
	"math"
)

// Solution from Project Euler itself

func SecondSolution(number int) int {
	lastFactor := 1
	if number%2 == 0 {
		lastFactor = 2
		number /= 2
		for number%2 == 0 {
			number /= 2
		}
	}

	factor := 3
	maxFactor := int(math.Sqrt(float64(number)))

	for number > 1 && factor <= maxFactor {
		if number%factor == 0 {
			number /= factor
			lastFactor = factor
			for number%factor == 0 {
				number /= factor
			}

			maxFactor = int(math.Sqrt(float64(number)))
		}
		factor += 2
	}

	if number == 1 {
		return lastFactor
	}

	return number
}
