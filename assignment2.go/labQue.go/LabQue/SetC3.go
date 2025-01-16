package main

import "fmt"
func calculate(a, b int) (int, int, int) {
    sum := a + b
    difference := a - b
    product := a * b

    return sum, difference, product
}

func main() {
    num1 := 10
    num2 := 5
    sum, diff, prod := calculate(num1, num2)
    fmt.Printf("Sum: %d, Difference: %d, Product: %d\n, sum, diff, prod)
}
