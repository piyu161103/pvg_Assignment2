package main
import "fmt"
func multiply(x,y,z int)int{
	return x*y*z
}
func main(){
	result:=multiply(3,5,6)
	fmt.Println("product",result)
}