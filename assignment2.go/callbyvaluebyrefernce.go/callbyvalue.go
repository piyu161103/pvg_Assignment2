package main
import "fmt"
func increment(num int){
	num = num + 1
	fmt.Println("Inside function:",num)
}
func main(){
	value:=10
	increment(value)
	fmt.Println("Outside function:",value)
}