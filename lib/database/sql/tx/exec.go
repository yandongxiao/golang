package main

import (
	"database/sql"
	"time"

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
	// guarantees that they’ll be executed on the same connection.
	// The methods on the Tx map one-for-one to methods you can call on the database itself, such as Query() and so forth.
	// NOTE: Prepared statements that are created in a transaction are bound exclusively to that transaction.
	// NOTE: While you are working inside a transaction you should be careful not to make calls to the Db variable.
	// Make all of your calls to the Tx variable that you created with db.Begin().
	// NOTE: If you need to work with multiple statements that modify connection state, you need a Tx even if you don’t want a transaction
	tx, err := db.Begin()
	checkError(err)

	// 在sleep期间，上面插入的数据不会被其它select语句感知到。
	// 我们在sql client上执行的一个select操作，也是一个事务。
	// 事务之间是隔离的
	_, err = tx.Exec("insert into user (age) values (?)", 10)
	checkError(err)
	time.Sleep(time.Second * 10)
	_, err = tx.Exec("insert into user (age) values (?)", 11)
	checkError(err)

	checkError(tx.Commit())
}
