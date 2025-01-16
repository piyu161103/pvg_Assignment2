package main

import "fmt"
func swap(a *int, b *int) {

    temp := *a
    *a = *b
    *b = temp
}

func main() {
    num1 := 10
    num2 := 20
    
    fmt.Println("Before swapping:")
    fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
    
    swap(&num1, &num2)
    
    fmt.Println("After swapping:")
    fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
}
