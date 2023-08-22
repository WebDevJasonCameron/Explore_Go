package main

import (
	"fmt"

	DbENV "book_archive/env"
)

func main() {
	fmt.Println(DbENV.DbNAME)
	fmt.Println(DbENV.DbHOST)
	fmt.Println(DbENV.DbNAME)
	fmt.Println(DbENV.DbPASSWORD)
	fmt.Println(DbENV.DbPORT)
	fmt.Println(DbENV.DbUSER)
}
