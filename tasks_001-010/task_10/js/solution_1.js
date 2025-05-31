/*
Сумма простых чисел меньше 10 равна 2 + 3 + 5 + 7 = 17.

Найдите сумму всех простых чисел меньше двух миллионов.
*/

function isPrime(num) {
	for (let i = 2; i <= Math.sqrt(num); i++)
		if (num % i === 0) return false
	return num > 1
}

const until = 2 * 10e6 //10
let result = 0

for (let i = 2; i < until; i++)
	if (isPrime(i)) result += i

console.log('Result:', result) // OUTPUT: 23514624000