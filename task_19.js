/*
Дана следующая информация (однако, вы можете проверить ее самостоятельно):

- 1 января 1900 года - понедельник.
- В апреле, июне, сентябре и ноябре 30 дней.
- В феврале 28 дней, в високосный год - 29.
- В остальных месяцах по 31 дню.
- Високосный год - любой год, делящийся нацело на 4, однако последний год века (ХХ00) является високосным в том и только том случае, если делится на 400.

Сколько воскресений выпадает на первое число месяца в двадцатом веке (с 1 января 1901 года до 31 декабря 2000 года)?
*/

const start = [1, 1, 1900]
const until = [31, 12, 2000]
const countFrom = 1901

let day = start[0],
	month = start[1],
	year = start[2],
	weekday = 1

let result = 0

function isLeapYear(year) {
	let century = year % 100 === 0 ? year / 100 : parseInt(year / 100 + 1)

	// Если год - последний год века
	if (year === century * 100) return year % 400 === 0
	else return year % 4 === 0
}

while (true) {
	// Смена дней недели
	if (weekday > 7) weekday = 1

	// Смена месяцев
	if (month === 4 || month === 6 || month === 9 || month === 11) {
		if (day > 30) {
			day = 1
			month++
		}
	}
	else if (month === 2) {
		if (isLeapYear(year)) {
			if (day > 29) {
				day = 1
				month++
			}
		}
		else {
			if (day > 28) {
				day = 1
				month++
			}
		}
	}
	else {
		if (day > 31) {
			day = 1
			month++
		}
	}

	// Смена года
	if (month > 12) {
		month = 1
		year++
	}

	if (year >= countFrom && day === 1 && weekday === 7) {
		console.log(`sunday on ${day}.${month}.${year}`)
		result++
	}
	if (day === until[0] && month === until[1] && year === until[2]) break

	day++
	weekday++
}

console.log('\nResult', result)