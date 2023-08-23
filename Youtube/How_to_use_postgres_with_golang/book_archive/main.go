package main

import (
	DbEnv "book_archive/env"
	"os"

	"github.com/jmoiron/sqlx"
)

// STRUCTs
type Person struct {
	Name  string
	Email string `gorm:"typevarchar(100);unique_index`
	Books []Book
}

type Book struct {
	Title      string
	Author     string
	CallNumber string `gorm:"unique_index`
	PersonID   int
}

func main() {

	// Loading ENV VARs
	//dbDialect := os.Getenv(DbEnv.DbDIALECT)
	//dbHost := os.Getenv(DbEnv.DbHOST)
	//dbPort := os.Getenv(DbEnv.DbPORT)
	dbUser := "user=" + DbEnv.DbUSER
	dbName := "name=" + DbEnv.DbNAME
	dbPassword := os.Getenv(DbEnv.DbPASSWORD)

	dbString := dbUser + " " + dbName + " " + "sslmode=disable"

	// Openning Connection to DB
	db, err := sqlx.Connect("postgres", dbString)

}
