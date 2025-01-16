package main

import "fmt"

func calculate(a int, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    
    num1 := 5
    num2 := 3
    
    sum, product := calculate(num1, num2)
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
