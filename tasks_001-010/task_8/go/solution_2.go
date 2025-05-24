package solution

func SecondSolution(adjacentDigits int) int {
	if adjacentDigits < 1 {
		return 0
	}

	var largestProduct int

	for i := 0; i < len(hugeNumber); {
		if i+adjacentDigits >= len(hugeNumber) {
			break
		}

		skipping := false
		product := 1

		for j := i; j < i+adjacentDigits; j++ {
			digit := int(hugeNumber[j] - '0') // byte - rune('0') -> int ---> 55 - 48 = 7

			// there's no chance it is going to be the largest product if it's multiplied by 0 or 1
			if digit == 0 || digit == 1 {
				i = j + 1
				skipping = true
				break
			}

			product *= digit
		}

		if !skipping {
			i++
		}

		if product > largestProduct {
			largestProduct = product
		}
	}

	return largestProduct
}
