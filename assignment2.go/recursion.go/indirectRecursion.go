package main

import "fmt"
func isEven(n int) bool {
    if n == 0 {
        return true 
    }
    return isOdd(n - 1) 
}

func isOdd(n int) bool {
    if n == 0 {
        return false 
    }
    return isEven(n - 1) 
}

func main() {
    number := 7
    if isEven(number) {
        fmt.Printf("%d is even.\n", number)
    } else {
        fmt.Printf("%d is odd.\n", number)
    }
}
