package main

import "fmt"

type A struct {
	Content string
}

//func (a A) Save() {
//	fmt.Println("Saving in the Terminal .... " + a.Content)
//}

type APersistence struct {
}

func (ap APersistence) Save(a A) {
	fmt.Println("Saving in the Terminal .... " + a.Content)
}

func main() {

	a := A{Content: "hello how are you"}

	//you can do this way
	APersistence{}.Save(a)

	//or

	ap := APersistence{}
	ap.Save(a)

	//a := A{Content:"hello how are you"}
	//a.Content
	//
	////a := A{Content:"hello how are you"}
	////a.Content

}
