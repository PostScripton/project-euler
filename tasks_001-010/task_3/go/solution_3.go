package solution

import (
	"math"
)

// Adjusted solution from Project Euler itself

func ThirdSolution(number int) int {
	var maxFactor int

	for i := 3; i < int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			number /= i
			maxFactor = i
		}
	}

	if number > maxFactor {
		maxFactor = number
	}

	return maxFactor
}
