package solution

func ThirdSolution(adjacentDigits int) int {
	if adjacentDigits < 1 {
		return 0
	}

	var largestProduct int

	for i := 0; i < len(hugeNumber); {
		if i+adjacentDigits >= len(hugeNumber) {
			break
		}

		skipping := false
		product := int(hugeNumber[i] - '0') // skipping first item in loop saves a lot of time

		for j := i + 1; j < i+adjacentDigits; j++ {
			digit := int(hugeNumber[j] - '0')

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
