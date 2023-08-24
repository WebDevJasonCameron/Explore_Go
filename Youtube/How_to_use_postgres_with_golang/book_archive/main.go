package main

import (
	DbEnv "book_archive/env"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// SCHEMA
var schema = `
CREATE TABLE person (
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT NOT NULL,
);
CREATE TABLE book (
	id SERIAL PRIMARY KEY,
	title TEXT,
	author TEXT,
	call_number TEXT,
);
CREATE TABLE person_to_book_ownership (
	person_id,
	book_id
);
`

// STRUCTs
type Person struct {
	id    uint32
	Name  string
	Email string
}

type Book struct {
	id         uint32
	Title      string
	Author     string
	CallNumber string
}

type PersonToBookOwnership struct {
	PersonID []Person
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
