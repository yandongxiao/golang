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
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(db.Ping())
	stmt, err := db.Prepare("insert into user (age) values (?)")
	checkError(err)
	defer stmt.Close() // Tx创建的stmt, 无需Close. Commit后会自动close

	result, err := stmt.Exec(100)
	checkError(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
