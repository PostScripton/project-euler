package main

import "fmt"

// Each third Fibonacci number is even, thus we can get rid of "is even" checks
// 1 1 (2) 3 5 (8) 13 21 (34) 55 89 (144) ...

const exceedNumber = 4000000

func main(){
    var result int

    first := 1
    second := 1
    third := first + second

    for third < exceedNumber {
        result += third

        first = second + third // 13 = 5 + 8
        second = third + first // 21 = 8 + 13
        third = first + second // 34 = 13 + 21
    }

    fmt.Println("Result:", result)
}
