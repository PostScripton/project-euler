package solution

func FirstSolution() int {
	var result int

	for ri, row := range grid {
		for ci, columnValue := range row {
			// Left to right: -
			if ci+length-1 < len(row) {
				product := columnValue

				for i := 1; i < length; i++ {
					product *= grid[ri][ci+i]
				}

				if product > result {
					result = product
				}
			}

			// Top to bottom: |
			if ri+length-1 < len(grid) {
				product := columnValue

				for i := 1; i < length; i++ {
					product *= grid[ri+i][ci]
				}

				if product > result {
					result = product
				}
			}

			// Forth diagonal: /
			if ri+length-1 < len(grid) && ci-length-1 >= 0 {
				product := columnValue

				for i := 1; i < length; i++ {
					product *= grid[ri+i][ci-i]
				}

				if product > result {
					result = product
				}
			}

			// Back diagonal: \
			if ri+length-1 < len(grid) && ci+length-1 < len(row) {
				product := columnValue

				for i := 1; i < length; i++ {
					product *= grid[ri+i][ci+i]
				}

				if product > result {
					result = product
				}
			}
		}
	}

	return result
}
