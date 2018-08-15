package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)
	fmt.Println(db.Ping())
}
