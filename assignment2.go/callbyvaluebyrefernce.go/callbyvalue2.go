package main

import "fmt"

func modifyByValue(x int) {
    x = 10
}

func main() {
    num := 5
    fmt.Println("Before modifyByValue:", num) 
    modifyByValue(num)
    fmt.Println("After modifyByValue:", num)  
}
