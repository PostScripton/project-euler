package main

import "fmt"

const until = 1000

func main(){
    var result int

    for num := 0; num < until; num++ {
        if num % 3 == 0 || num % 5 == 00 {
            result += num
        }
    }
	
    fmt.Println("Result:", result)
}
