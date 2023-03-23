# SOLID-PRINCIPLES-Golang
- The SOLID principles are a set of guidelines for writing maintainable and scalable software.
- Design Go programs that are well engineered, decoupled, reusable, and responsive to changing requirements.
- Avoid writing bad code or bad design, which has these characteristics: rigidity, fragility, immobility, complexity, verbosity.


SOLID stands for:

- S - Single-responsiblity Principle
- O - Open-closed Principle
- L - Liskov Substitution Principle
- I - Interface Segregation Principle
- D - Dependency Inversion Principle

## Single Responsibility Principle (SRP)

The Single Responsibility Principle states that a class should have only one reason to change. In other words, a class should only have one responsibility. In Golang, this can be achieved by creating small, focused functions that perform specific tasks.

“One function or type should have one and only one job, and one and only one responsibility”

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


## Open/Closed Principle (OCP)

The Open/Closed Principle states that a class should be open for extension but closed for modification. In other words, the behavior of a class should be able to be extended without modifying its source code. In Golang,this can be achieved by using interfaces and polymorphism.

It imply 2 things :
- be able to override a struct
    type A struct {
        year int
    }
    func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

    type B struct {
        A
    }

    func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year) }

- Use the Strategy Design Pattern


## Liskov Substitution Principle
The Liskov Substitution Principle states that objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program. In other words, a derived class should be able to substitute its base class without affecting the functionality of the program.
Here’s an example of how the Liskov Substitution Principle can be applied in Golang:


```package main

import "fmt"

// This is an interface that defines a shape.

type Shape interface {
    Area() float64
}

// This is a struct that implements the Shape interface.

type Rectangle struct {
    width float64
    height float64
}

// This method calculates the area of a rectangle.

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// This is a struct that implements the Shape interface.

type Square struct {
    side float64
}

// This method calculates the area of a square.

func (s Square) Area() float64 {
    return s.side * s.side
}

// This function takes a Shape as an argument and calculates its area.

func calculateArea(s Shape) float64 {
    return s.Area()
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Println("Area of rectangle:", calculateArea(rect))

    square := Square{side: 5}
    fmt.Println("Area of square:", calculateArea(square))
}```

In this example, the Square struct implements the Shape interface and thus can be used wherever a Shape is expected. This allows us to substitute a Square object for a Rectangle object without affecting the correctness of the program, as both have a Area() method that implements the Shape interface. This is an example of the Liskov Substitution Principle in action.

## Interface Segregation Principle (ISP)
The Interface Segregation Principle states that clients should not be forced to depend on interfaces they do not use. In Golang, this can be achieved by creating smaller, focused interfaces that provide only the functionality needed by a specific client.

Clients should not be forced to depend on methods they do not use.

It means that we shouldn’t create big struct with a lot of behaviors, we should isolate behaviors.

As example, we could think in a Vehicle with almost 2 methods :

type Vehicle struct {
 
}
func (v Vehicle) Accelerate() {
     fmt.Println("Accelerating ....")
}
func (v Vehicle) PlayCD() {
     fmt.Println("Playnig Guns n Roses")
}
A client using our Vehicle, wants to “extend” our vehicle to implement a Bus struct (in composition), but his Bus does not need to PlayCD. But he will be obliged to have the method, even if he will not use it :S

So on this case we should separate in 2 different struct the behaviors (in SRP) CDPlayer and Accelerator.define 2 interfaces, and 2 different struct.


type VehicleCDPlayer interface {
    PlayCD()
}
type VehicleCdPlay struct {
}
func (v VehicleCdPlay) PlayCD() {
     fmt.Println("Playnig Guns n Roses")
}
type VehicleAccelerator interface {
    Accelerate()
}
type VehicleAccelerate struct {
 
}
func (v VehicleAccelerate) Accelerate() {
     fmt.Println("Accelerating ....")
}
type Vehicle struct {
 VehicleAccelerate
 VehicleCdPlay
}



## Dependency Inversion Principle (DIP)
The Dependency Inversion Principle (DIP) states that high-level modules should not depend on low-level modules, but both should depend on abstractions.

Here’s an example of how the Dependency Inversion Principle can be applied in Golang:

package main

import "fmt"

// This is an interface that defines basic operations for a database.
type Database interface {
    Connect()
    Store(data string)
}

// This is a struct that implements the Database interface.
type MySQL struct {}

// This method connects to a MySQL database.
func (m MySQL) Connect() {
    fmt.Println("Connecting to MySQL database...")
}

// This method stores data in a MySQL database.
func (m MySQL) Store(data string) {
    fmt.Println("Storing data in MySQL database:", data)
}

// This is a struct that implements the Database interface.
type PostgreSQL struct {}

// This method connects to a PostgreSQL database.
func (p PostgreSQL) Connect() {
    fmt.Println("Connecting to PostgreSQL database...")
}

// This method stores data in a PostgreSQL database.
func (p PostgreSQL) Store(data string) {
    fmt.Println("Storing data in PostgreSQL database:", data)
}

// This is a struct that depends on a Database interface.
type Service struct {
    db Database
}

// This method sets the Database for the Service.
func (s *Service) SetDatabase(db Database) {
    s.db = db
}

// This method stores data in a database using the Database interface.
func (s *Service) StoreData(data string) {
    s.db.Connect()
    s.db.Store(data)
}

func main() {
    // This creates a Service that uses a MySQL database.
    mysqlService := Service{db: MySQL{}}
    mysqlService.StoreData("Hello, world!")

    // This creates a Service that uses a PostgreSQL database.
    postgresService := Service{db: PostgreSQL{}}
    postgresService.StoreData("Hello, world!")
}


In this example, the Service struct depends on the Database interface, and the MySQL and PostgreSQL structs implement the Database interface. This means that the Service struct can store data in either a MySQL or PostgreSQL database, without having to know or care which database it’s using. This is an example of the Dependency Inversion Principle in action. The high-level Service module depends on an abstraction (the Database interface), while both the low-level MySQL and PostgreSQL modules depend on the same abstraction. This allows the Service module to be decoupled from the specific implementation of the database, making it more flexible and maintainable.                            
