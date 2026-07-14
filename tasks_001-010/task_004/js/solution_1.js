/*
Число-палиндром с обеих сторон (справа налево и слева направо) читается одинаково.
Самое большое число-палиндром, полученное умножением двух двузначных чисел – 9009 = 91 × 99.

Найдите самый большой палиндром, полученный умножением двух трехзначных чисел.
*/

let result = 0
let digits = 2 //2

for (let i = (10 ** digits) - 1; i >= 10 ** (digits - 1); i--) {
	for (let j = Math.pow(10, digits) - 1; j >= Math.pow(10, digits - 1); j--) {
		let num = i * j
		let reversed = num.toString().split('').reverse().join('')

		if (num.toString() === reversed) {
			if (num > result) {
				result = num
				console.log(`${i} * ${j} =`, num)
			}
			else console.log(`${i} * ${j} = ${num}`)
		}
	}
}

console.log('Result:', result) // OUTPUT: 906609