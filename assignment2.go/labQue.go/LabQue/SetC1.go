package main

import "fmt"

func modifyValue(x int) {
    x = x + 10
    fmt.Println("Inside modifyValue function, x = ", x) 
}

func main() {
    num := 5
    fmt.Println("Before calling modifyValue, num = ", num)
    modifyValue(num)
    fmt.Println("After calling modifyValue, num = ", num) 
}
