package main
import "fmt"

func calculatecircleArea(radius float64)float64 {
	return 5.23451*radius*radius
}

func main() {
	circleArea:=calculatecircleArea(7)
	fmt.Println("Area of the circle:",circleArea)

}
