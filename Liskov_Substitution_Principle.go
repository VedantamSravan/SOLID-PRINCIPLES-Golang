package main

import "fmt"

type transport interface {
	getName() string
}
type vehicle struct {
	name string
}

func (v vehicle) getName() string {
	return v.name
}

type car struct {
	vehicle
	wheel int
	doors int
}
type motorcyle struct {
	vehicle
	wheel int
}
type printer struct {
}

func (printer) printTransportName(p transport) {
	fmt.Println("Name:", p.getName())
}
func main() {
	v := vehicle{name: "ford"}
	c := car{
		vehicle: vehicle{name: "Car-name"},
		wheel:   4,
		doors:   2,
	}
	m := motorcyle{
		vehicle: vehicle{name: "motorcycle-name"},
		wheel:   2,
	}
	pr := printer{}
	pr.printTransportName(v)
	pr.printTransportName(c)
	pr.printTransportName(m)
}

//type A struct {
//
//}
////by using interface
//type tester interface {
//	Test()
//}
//
//func (a A) Test() {
//	fmt.Println("Printing A")
//}
//type B struct {
//	A
//}
//// PANIC : cannot use a (type B) as type A in argument to ImpossibleLiskovSubstitution
////func ImpossibleLiskovSubstitution(a A) {
////	a.Test()
////}
//
//func PossibleLiskovSubstitution(a tester) {
//	a.Test()
//}
//
//func main() {
//	a := B{}
//	PossibleLiskovSubstitution(a)
//
//
//
//}
