package solution

func SecondSolution(until int) int {
	result := 5 // we know that 2 and 3 are prime

	for number := 5; number <= until; number += 2 { // all next prime numbers are odd, then add 2
		if isPrimeNumber(number) {
			result += number
		}
	}

	return result
}
