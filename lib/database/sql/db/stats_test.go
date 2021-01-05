package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ExampleStats() {
	// panic: sql: unknown driver "mysql" (forgotten import?)
	// return db handle
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)
	stats := db.Stats()
	fmt.Println(db.Ping())
	// 参见Readme.md中关于Open Conn 和 Idle Conn的区别
	fmt.Println(stats.OpenConnections)

	// Output:
	// <nil>
	// 0
}
