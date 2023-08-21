package main

import (
	// "fmt"
	// "os"

	"fmt"

	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber string `gorm:"unique_index"`
	PersonID   string
}

var db *gorm.DB
var err error

func main() {

	greeting := modularTest()

	fmt.Println(greeting)

	// Loading evn vars
	// dialect := os.Getenv(DIALECT)
	// host := os.Getenv(HOST)
	// dbPort := os.Getenv(DBPORT)
	// user := os.Getenv(USER)
	// dbName := os.Getenv(NAME)
	// password := os.Getenv(PASSWORD)

	// DB connection string
	// dbURI := fmt.Sprintf("host=%s user=%s ssllmode=disable password%s port=%s", host, user, password, dbPort)
}
