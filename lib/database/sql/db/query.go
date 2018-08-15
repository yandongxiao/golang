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

func query() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)

	rows, err := db.Query("select id, age from user where age=?", 10)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var age int
		checkError(rows.Scan(&id, &age))
		fmt.Println(id, age)
	}
	checkError(rows.Err())
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)

	// 如果有多行数据满足要求，则只返回其中一条
	// If an error occurs during the execution of the statement, that error
	// will be returned by a call to Scan on the returned *Row 解释QueryRow没有返回error
	row := db.QueryRow("select id, age from user where age=?", 10)
	var id, age int
	checkError(row.Scan(&id, &age))
	fmt.Println(id, age)
}
