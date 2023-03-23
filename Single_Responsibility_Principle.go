package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}
type square struct {
	length float32
}

//here this area function is doing two things
//one is calculating area and also printing it
//so what it should be doing calculating area and some body else should do prinitng.

/*func (c circle)area(){

	fmt.Println("Area of circle is",math.Pi*c.radius*c.radius)
}*/

//so here we applied sinlge Responsibility Principle

//separating print() and cliculate() in different functions as below

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func printArea(value float32) {
	fmt.Println("Area of circle is", value)

}
func main() {

	circle := circle{radius: 1}
	//circle.area()
	printArea(circle.area())
}
