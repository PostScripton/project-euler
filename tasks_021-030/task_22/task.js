/*
Используйте names.txt, текстовый файл размером 46 КБ, содержащий более пяти тысяч имен.
Начните с сортировки в алфавитном порядке. Затем подсчитайте алфавитные значения каждого имени и умножьте это значение на порядковый номер имени
в отсортированном списке для получения количества очков имени.

Например, если список отсортирован по алфавиту, имя COLIN (алфавитное значение которого 3 + 15 + 12 + 9 + 14 = 53) является 938-ым в списке.
Поэтому, имя COLIN получает 938 × 53 = 49714 очков.

Какова сумма очков имен в файле?
*/

const fs = require('fs')

function get_names() {
	// array of names
	return fs.readFileSync(__dirname + '/names.txt').toString().split(',').map(v => v.slice(1, v.length -1))
}

const letterOrder = letter => 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.indexOf(letter.toUpperCase()) + 1

let result = 0

get_names().sort().forEach((name, index) => {
	let alphabeticValue = name.split('').reduce((acc, cur) => acc + letterOrder(cur), 0)
	result += alphabeticValue * (index + 1)
})

console.log('Result:', result) //OUTPUT: 871198282