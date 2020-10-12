/*
2520 - самое маленькое число, которое делится без остатка на все числа от 1 до 10.

Какое самое маленькое число делится нацело на все числа от 1 до 20?
*/

let num = 1
let until = 20 //10

function smallestMultiple() {
	while (true) {
		let smallest = true
		for (let i = 2; i <= until; i++) {
			if (num % i !== 0) {
				smallest = false
				break
			}
		}
		if (smallest) return num
		num++
	}
}

console.log(`Result: ${smallestMultiple()}`) // OUTPUT: 232792560