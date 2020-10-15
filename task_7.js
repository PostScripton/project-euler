/*
Выписав первые шесть простых чисел, получим 2, 3, 5, 7, 11 и 13. Очевидно, что 6-ое простое число - 13.

Какое число является 10001-ым простым числом?
*/

function isPrime(num) {
	for (let i = 2; i <= Math.sqrt(num); i++)
		if (num % i === 0) return false
	return num > 1
}

const until = 10001 //6
let num = 2
let result = 0
while(true) {
	if (isPrime(num)) {
		if (++result === until) {
			result = num
			break
		}
	}
	num++
}

console.log('Result:', result) // OUTPUT: 104743