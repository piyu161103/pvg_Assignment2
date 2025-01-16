package main

import "fmt"

func modifyByReference(x *int) {
    *x = 10
}

func main() {
    num := 7
    fmt.Println("Before modifyByReference:", num) 
    modifyByReference(&num) 
    fmt.Println("After modifyByReference:", num)  
}
