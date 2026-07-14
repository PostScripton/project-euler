package solution

func FirstSolution(until int) int {
	for number := until; ; number++ {
		ok := true

		for i := 2; i <= until; i++ {
			if number%i != 0 {
				ok = false
				break
			}
		}

		if ok {
			return number
		}
	}
}
