package main

import "fmt"

func calculate(a int, b int) (sum int, difference int) {
    sum = a + b
    difference = a - b
    return
}

func main() {
    sum, difference := calculate(10, 5)
    fmt.Printf("Sum: %d, Difference: %d\n", sum, difference)
}
