package main
import "fmt"
func square(x int)int{
	return x*x
}
func main(){
	num:=6
	result:=square(num)
	fmt.Println("sqaure:",result)
}