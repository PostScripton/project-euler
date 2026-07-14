package solution

import (
	"math"
)

func SecondSolution(digits int) int {
	var result int

	startingNumber := int(math.Pow(10, float64(digits))) - 1
	endingNumber := int(math.Pow(10, float64(digits-1)))

	for i := startingNumber; i >= endingNumber; i-- {
		// we don't want to multiply the same numbers as i=99,j=92 and i=92,j=99
		// thus, i <= j always
		for j := startingNumber; j >= i; j-- {
			product := i * j

			if isPalindrome(product) && product > result {
				result = product
			}
		}
	}

	return result
}
