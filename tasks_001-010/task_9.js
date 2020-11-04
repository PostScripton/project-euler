/*
Тройка Пифагора - три натуральных числа a < b < c, для которых выполняется равенство

a^2 + b^2 = c^2
Например, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

Существует только одна тройка Пифагора, для которой a + b + c = 1000.
Найдите произведение abc.
*/

const perimeter = 1000
let result = 0

for (let a = 1; a <= perimeter + 1 && result === 0; a++){
	for (let b = a + 1; b <= perimeter + 1 && result === 0; b++) {
		let c = perimeter - a - b
		if (a ** 2 + b ** 2 === c ** 2) {
			console.log(`${a} + ${b} + ${c} = ${a + b + c}`)
			result = a * b * c
		}
	}
}

console.log('Result:', result) // OUTPUT: 31875000