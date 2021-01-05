package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)

	// 1. you should use question mark parameters method
	// 2. do not concatenate strings together. 这样的话，就起不到防SQL注入了
	// 3. db.Query并没有返回任何数据，just a cursor.
	// 4. If a function name includes Query, it is designed to ask a question of the database,
	//    and will return a set of rows, even if it’s empty. NOTE: Statements that don’t return rows
	//    should not use Query functions; they should use Exec()
	// Under the hood, db.Query() actually prepares, executes, and closes a prepared statement.
	rows, err := db.Query("select id, age from user where age=?", 10)
	checkError(err)
	// You should always defer rows.Close())
	// NOTE:内部有一个到数据库的连接，如果不关闭则会造成资源的泄露
	// If rows.Close() returns an error, it’s unclear what you should do.
	// Logging the error message or panicing might be the only sensible thing,
	// and if that’s not sensible, then perhaps you should just ignore the error.)
	defer rows.Close()

	for rows.Next() {
		var id int
		var age int
		// Don’t just assume that the loop iterates until you’ve processed all the rows.
		checkError(rows.Scan(&id, &age))
		fmt.Println(id, age)
	}

	if err = rows.Err(); err != nil {
		rows.Close() // NOTE: 幂等操作，与之前的Close操作不冲突
		panic(err)
	}
}

func ExampleQuery() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)

	// 如果有多行数据满足要求，则只返回其中一条
	// If an error occurs during the execution of the statement, that error
	// will be returned by a call to Scan on the returned *Row 解释QueryRow没有返回error
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	var id, age int
	err = db.QueryRow("select id, age from user where age=?", 10).Scan(&id, &age)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			// An empty result is often not considered an error by application code,
			// and if you don’t check whether an error is this special constant,
			// you’ll cause application-code errors you didn’t expect
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(id, age)

	// Output:
	// 0 0
}
