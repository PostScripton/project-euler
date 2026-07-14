package solution

import (
	"math"
	"strconv"
)

func FirstSolution(digits int) int {
	var result int

	startingNumber := int(math.Pow(10, float64(digits))) - 1 // 999
	endingNumber := int(math.Pow(10, float64(digits-1)))     // 100

	for i := startingNumber; i >= endingNumber; i-- {
		for j := startingNumber; j >= endingNumber; j-- {
			product := i * j

			if isPalindrome(product) && product > result {
				result = product
			}
		}
	}

	return result
}

func isPalindrome(number int) bool {
	s := strconv.Itoa(number)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
