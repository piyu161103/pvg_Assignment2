package main

import "fmt"
func calculate(a int, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    sum, product := calculate(4, 5)
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
