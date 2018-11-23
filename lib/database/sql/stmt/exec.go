// NOTE: Avoiding Prepared Statements
//	1. The database doesn’t support prepared statements.
//  2. The statements aren’t reused enough to make them worthwhile, and
//     security issues are handled in other ways, so performance overhead is undesired
//  3. When you operate on a Tx object, your actions map directly to the one and only one connection underlying it.
//     the earlier cautions about repreparing do not apply.
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

	// 1. When you prepare a statement, it’s prepared on a connection in the pool.
	// 2. The Stmt object remembers which connection was used.
	// 3. When you execute the Stmt, it tries to use the connection.
	//    If it’s not available because it’s closed or busy doing something else,
	//    it gets another connection from the pool and re-prepares the statement
	//    with the database on another connection.
	stmt, err := db.Prepare("insert into user (age) values (?)")
	checkError(err)
	defer stmt.Close() // Tx创建的stmt, 无需Close. Commit后会自动close

	result, err := stmt.Exec(100)
	checkError(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
