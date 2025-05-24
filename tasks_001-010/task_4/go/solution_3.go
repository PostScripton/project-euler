package solution

import (
	"math"
)

func ThirdSolution(digits int) int {
	var result int

	startingNumber := int(math.Pow(10, float64(digits))) - 1
	endingNumber := int(math.Pow(10, float64(digits-1)))

	for i := startingNumber; i >= endingNumber; i-- {
		for j := startingNumber; j >= i; j-- {
			product := i * j

			// as all next sequences are only decreasing, then
			// there's no need to continue after finding palindrome
			if product <= result {
				break
			}

			if isPalindrome(product) && product > result {
				result = product
			}
		}
	}

	return result
}
