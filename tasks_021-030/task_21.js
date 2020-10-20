/*
Пусть d(n) определяется как сумма делителей n (числа меньше n, делящие n нацело).
Если d(a) = b и d(b) = a, где a ≠ b, то a и b называются дружественной парой, а каждое из чисел a и b - дружественным числом.

Например, делителями числа 220 являются 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 и 110, поэтому d(220) = 284. Делители 284 - 1, 2, 4, 71, 142, поэтому d(284) = 220.

Подсчитайте сумму всех дружественных чисел меньше 10000.
*/

function sumOfDivisors(n) {
	let sqrt = Math.sqrt(n)
	let sum = 1
	for (let i = 2; i <= sqrt; i++) {
		if (n % i === 0) {
			sum += i + (n / i)
			if (Number.isInteger(sqrt) && i === sqrt) sum -= i
		}
	}

	return sum
}

const until = 10000
let num = 1,
	result = 0

while (num < until) {
	let a = sumOfDivisors(num) //d(220) = 284
	let b = sumOfDivisors(a) //d(284) = 220
	if (num === b && num !== a) {
		console.log(`n: ${num} - a: ${a}, b: ${b}`)
		result += num
	}
	num++
}

console.log('\nResult:', result)