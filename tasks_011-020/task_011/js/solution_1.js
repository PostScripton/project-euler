const fs = require('fs')
const path = require('path')

const table = fs.readFileSync(path.join(__dirname, '..', 'grid.txt'), 'utf8')
	.trim()
	.split('\n')
	.map(row => row.split(' '))
const numsCount = 4

let result = 0

for (let row = 0; row < table.length; row++) {
	for (let column = 0; column < table[row].length; column++) {

		// Back diag \
		if (table[row + numsCount - 1] !== undefined) {
			if (table[row + numsCount - 1][column + numsCount - 1] !== undefined) {
				let product = table[row][column]
				for (let num = 1; num < numsCount; num++)
					product *= table[row + num][column + num]
				if (product > result) result = product
			}
		}

		// Forth diag /
		if (table[row + numsCount - 1] !== undefined) {
			if (table[row + numsCount - 1][column - numsCount + 1] !== undefined) {
				let product = table[row][column]
				for (let num = 1; num < numsCount; num++)
					product *= table[row + num][column - num]
				if (product > result) result = product
			}
		}

		// Top down |
		if (table[row + numsCount - 1] !== undefined) {
			let product = table[row][column]
			for (let num = 1; num < numsCount; num++)
				product *= table[row + num][column]
			if (product > result) result = product
		}

		// Left right —
		if (table[row][column + numsCount - 1] !== undefined) {
			let product = table[row][column]
			for (let num = 1; num < numsCount; num++)
				product *= table[row][column + num]
			if (product > result) result = product
		}
	}
}

console.log('Result:', result) // OUTPUT: 70600674
