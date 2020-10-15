/*
Если записать числа от 1 до 5 английскими словами (one, two, three, four, five), то используется всего 3 + 3 + 5 + 4 + 4 = 19 букв.

Сколько букв понадобится для записи всех чисел от 1 до 1000 (one thousand) включительно?


Примечание: Не считайте пробелы и дефисы. Например, число 342 (three hundred and forty-two) состоит из 23 букв, число 115 (one hundred and fifteen) - из 20 букв.
Использование "and" при записи чисел соответствует правилам британского английского.
*/

function inWords(num) {
	let str = num.toString()
	let words = [
		[1, 'one'],
		[2, 'two'],
		[3, 'three'],
		[4, 'four'],
		[5, 'five'],
		[6, 'six'],
		[7, 'seven'],
		[8, 'eight'],
		[9, 'nine'],
		[10, 'ten'],
		[11, 'eleven'],
		[12, 'twelve'],
		[13, 'thirteen'],
		[14, 'fourteen'],
		[15, 'fifteen'],
		[16, 'sixteen'],
		[17, 'seventeen'],
		[18, 'eighteen'],
		[19, 'nineteen'],
		[20, 'twenty'],
		[30, 'thirty'],
		[40, 'forty'],
		[50, 'fifty'],
		[60, 'sixty'],
		[70, 'seventy'],
		[80, 'eighty'],
		[90, 'ninety'],
	]

	if (words.some(v => v[0] === num)) return words.find(v => v[0] === num)[1]

	if (str.length === 2) {
		return inWords(parseInt(str[0] + '0')) + '-' + inWords(parseInt(str[1]))
	}
	else if (str.length === 3) {
		if (num % 100 === 0) return inWords(num / 100) + ' hundred'
		return inWords(parseInt(str[0] + '00')) + ' and ' + inWords(parseInt(str.slice(1, 3)))
	}
	else if (str.length === 4) {
		if (num % 1000 === 0) return inWords(num / 1000) + ' thousand'
	}

	return 'Big enough'
}

const until = 1000 //5
const arr = [...Array(until).keys()].map(v => ++v)

let result = arr.reduce((acc, cur) => {
	console.log(`${cur} - ${inWords(cur)}`)
	return acc + inWords(cur).replace('-', ' ').split(' ').join('').length
}, 0)

console.log('Result:', result)