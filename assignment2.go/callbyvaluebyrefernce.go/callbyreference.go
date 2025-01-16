package main
import "fmt"
func incrementbyReference(num*int){
	*num = *num + 1
	fmt.Println("Inside function:",*num)
}
func main(){
	value:=10
	incrementbyReference(&value)
	fmt.Println("Outside function:",value)
}