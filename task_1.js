/*
Если выписать все натуральные числа меньше 10, кратные 3 или 5, то получим 3, 5, 6 и 9. Сумма этих чисел равна 23.

Найдите сумму всех чисел меньше 1000, кратных 3 или 5.
*/

const until = 1000 //10
let arr = Array.from(Array(until - 1).keys()).map(v => v + 1) // массив от 1 до 999

let result = arr.reduce((acc, cur) => (cur % 3 === 0 || cur % 5 === 0) ? acc + cur : acc, 0)
console.log(`Result: ${result}`)