package main

import (
	DbEnv "book_archive/env"
	"fmt"
	"os"

	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
)

// STRUCTs
type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber string `gorm:"unique_index`
	PersonID   int
}

// VARs
var db *gorm.DB
var err error

func main() {

	// Loading ENV VARs
	dbDialect := os.Getenv(DbEnv.DbDIALECT)
	dbHost := os.Getenv(DbEnv.DbHOST)
	dbPort := os.Getenv(DbEnv.DbPORT)
	dbUser := os.Getenv(DbEnv.DbUSER)
	dbName := os.Getenv(DbEnv.DbNAME)
	dbPassword := os.Getenv(DbEnv.DbPASSWORD)

	// DB Connection String
	dbUri := fmt.Sprintf("host=%s user%s dbname%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPassword, dbPort)

	// Openning Connection to DB
	db, err = gorm.Open(dbDialect, dbUri)

}
