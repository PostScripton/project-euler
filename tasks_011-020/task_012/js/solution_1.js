const until = 500
let num = 0
let triangle = 0
let divisor = 0

while (true) {
	divisor = 2 // 1 и само число
	triangle += num
	for (let i = 2; i <= Math.sqrt(triangle); i++) {
		if (triangle % i === 0) {
			divisor += 2
			if (Number.isInteger(Math.sqrt(triangle))) divisor--
		}
	}
	if (divisor > until) break
	num++
}