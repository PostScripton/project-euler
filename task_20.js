/*
n! означает n × (n − 1) × ... × 3 × 2 × 1

Например, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
и сумма цифр в числе 10! равна 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Найдите сумму цифр в числе 100!.
*/

const until = 100
const factorial = [...Array(until).keys()].map(v => ++v).reverse()

let product = BigInt(factorial.reduce((acc, cur) => BigInt(acc) * BigInt(cur)))
let result = product.toString().split('').map(v => parseInt(v)).reduce((acc, cur) => acc + cur)

console.log('Result:', result)