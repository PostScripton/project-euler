/*
Простые делители числа 13195 - это 5, 7, 13 и 29.

Каков самый большой делитель числа 600851475143, являющийся простым числом?
*/

let num = 600851475143 //13195

let max = 0
for (let i = 3; i < Math.round(Math.sqrt(num)); i += 2) {
	if (num % i === 0) {
		console.log(`num(${num}) / prime(${i}) = ${num / i}`)
		num /= i
		max = i
	}
}
if (num > max) max = num

console.log('Result:', max) // OUTPUT: 6857
