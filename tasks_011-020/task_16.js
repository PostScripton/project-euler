/*
2^15 = 32768, сумма цифр этого числа равна 3 + 2 + 7 + 6 + 8 = 26.

Какова сумма цифр числа 2^1000?
*/

const power = 1000 //15
let result = BigInt(Math.pow(2, power)).toString().split('').reduce((acc, cur) => acc + parseInt(cur), 0)

console.log('Result:', result) // OUTPUT: 1366