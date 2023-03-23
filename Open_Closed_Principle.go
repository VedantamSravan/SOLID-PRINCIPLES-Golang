package main

import (
	"fmt"
)

//If we want to add a new Behavior (the minus for example) it would not be possible.

//So we will use the strategy Pattern.
/*type Calcul struct {
}
func (c Calcul) Calculate (a int, b int) int {
	return a + b
}*/

type Calculer interface {
	Execute(int, int) int
}
type Add struct {
}

func (add Add) Execute(a int, b int) int {
	return a + b
}

type Minus struct {
}

func (minus Minus) Execute(a int, b int) int {
	return a - b
}

type Calcul struct {
	cel Calculer
}

func (c Calcul) Calculate(a int, b int) int {
	return c.cel.Execute(a, b)
}
func main() {

	a := Calcul{cel: Add{}}
	b := Calcul{cel: Minus{}}
	fmt.Println(a.Calculate(3, 1)) // 4
	fmt.Println(b.Calculate(3, 1)) // 2
	//c := Calcul{}
	//val1:= c.Calculate(10,20)
	//fmt.Println(val1)

}
