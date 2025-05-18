package main

import "fmt"

const exceedNumber = 4000000

func main(){
    var result int

    first := 1
    second := 2

    for second < exceedNumber {
        if second % 2 == 0 {
            result += second
        }

        next := first + second
        first = second
        second = next
    }
	
    fmt.Println("Result:", result)
}
