/*
Найдите первые десять цифр суммы следующих ста 50-значных чисел: (number.txt)
*/

const fs = require('fs')

function get_numbers() {
	// array of string numbers
	return fs.readFileSync(__dirname + '/number.txt').toString().split('\r\n')
}

let result = get_numbers().reduce((acc, cur) => BigInt(acc) + BigInt(cur)).toString().slice(0, 10)

console.log(`Result: ${result}`) // OUTPUT: 5537376230
