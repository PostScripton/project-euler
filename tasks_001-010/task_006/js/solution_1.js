/*
Сумма квадратов первых десяти натуральных чисел равна
1^2 + 2^2 + ... + 10^2 = 385

Квадрат суммы первых десяти натуральных чисел равен
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Следовательно, разность между суммой квадратов и квадратом суммы первых десяти натуральных чисел составляет 3025 − 385 = 2640.

Найдите разность между суммой квадратов и квадратом суммы первых ста натуральных чисел.
*/

let until = 100 //10
let nums = [...Array(until).keys()].map(v => v + 1) // массив от 1 до 10

let sumOfSquares = nums.reduce((acc, cur) => acc + Math.pow(cur, 2))
let squareOfSum = Math.pow(nums.reduce((acc, cur) => acc + cur), 2)

let result = squareOfSum - sumOfSquares

console.log('Result:', result) // OUTPUT: 25164150
