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

func ExampleExec() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)

	// Exec executes a query without returning any rows.
	// NOTE: use an exec for anything you does not return rows
	// 虽然可以使用db.Query, 但是它非常容易写成_, err := db.Query("xxx").
	// 这时由于忽略了第一个参数，导致rows.Close()没有调用，造成资源泄露
	result, err := db.Exec("insert into user (name, age) values('jack',?)", 999)
	checkError(err)

	// fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// Output:
	// 1 <nil>
}
