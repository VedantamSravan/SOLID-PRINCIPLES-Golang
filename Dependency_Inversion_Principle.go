package main

import "fmt"

// This is an interface that defines basic operations for a database.
type Database interface {
	Connect()
	Store(data string)
}

// This is a struct that implements the Database interface.
type MySQL struct{}

// This method connects to a MySQL database.
func (m MySQL) Connect() {
	fmt.Println("Connecting to MySQL database...")
}

// This method stores data in a MySQL database.
func (m MySQL) Store(data string) {
	fmt.Println("Storing data in MySQL database:", data)
}

// This is a struct that implements the Database interface.
type PostgreSQL struct{}

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
