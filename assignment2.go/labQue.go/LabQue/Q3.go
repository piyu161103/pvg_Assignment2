package main

import "fmt"

func isPalindrome(n int) bool {
    original := n
    reversed := 0

    for n > 0 {
        remainder := n % 10
        reversed = reversed*10 + remainder
        n = n / 10
    }
    return original == reversed
}

func main() {
    var number int
    fmt.Print("Enter a number: ")
    fmt.Scan(&number)

    if isPalindrome(number) {
        fmt.Printf("%d is a palindrome.\n", number)
    } else {
        fmt.Printf("%d is not a palindrome.\n", number)
    }
}
