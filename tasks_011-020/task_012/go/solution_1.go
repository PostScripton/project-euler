package solution

import (
	"math"
)

func FirstSolution(divisors int) int {
	var triangle int

	for n := 1; ; n++ {
		triangle += n

		if getDivisorsCount(triangle) > divisors {
			break
		}
	}

	return triangle
}

func getDivisorsCount(n int) int {
	var result int

	// d <= Sqrt(n) is equivalent to d*d <= n
	//
	// For a number we don't need to check big divisors, we can check only small ones and calculate big ones on fly:
	// E.g. for 28, the divisors will be: 1, 2, 4 | 7, 14, 28
	// Sqrt(28) ~= 5.29 but we don't need to be that precise,
	// the root just says until which number we should check divisors.
	// To find big divisors, we may take the small one and divide the number by it:
	//   28 / 1 = 28
	//   28 / 2 = 14
	//   28 / 4 = 7
	for d := 1; float64(d) <= math.Sqrt(float64(n)); d++ {
		// The divisor of N is the number by which N is divisible without remainder
		if n%d != 0 {
			continue
		}

		oppositeDivisor := n / d

		if oppositeDivisor == d {
			// For a perfect square, d and n/d are the same divisor,
			// so it must be counted only once
			result += 1
		} else {
			// Otherwise, it counts as 2 because there are 2 different divisors
			// (small one (before square root) and big one (after square root))
			result += 2
		}
	}

	return result
}
