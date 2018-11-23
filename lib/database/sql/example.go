// NOTE: you can not use uint64 with high bit set in parameters.
// NOTE : package database/sql 包里面有数据类型DB, Stmt, Tx 三种类型，
//        他们都支持Exec、 Query、QueryRow操作
// 1. do not defer in long running functions
// 2. Do not use connection state (use xx-db, set utf-8). Use Tx to bind to a connection
// 3. DB is not a connection. database/sql对上屏蔽了连接的概念，所以设置连接的属性或状态时无效的
// 4. Do not use BEGIN and COMMIT via SQL. (use Tx)
//
// NOTE: Avoiding Prepared Statements. 注意不推荐使用db.Prepare方法
//	1. The database doesn’t support prepared statements.
//  2. The statements aren’t reused enough to make them worthwhile, and
//     security issues are handled in other ways, so performance overhead is undesired
//
// NOTE: What if your connection to the database is dropped, killed, or has an error?
// You don’t need to implement any logic to retry failed statements when this happens.
// As part of the connection pooling in database/sql, handling failed connections is built-in. 10次重试
//
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 1. a sql.DB isn’t a database connection, represents your database
	// 2. It’s an abstraction of the interface and existence of a database,
	//    which might be as varied as a local file, accessed through a network
	//    connection, or in-memory and in-process.
	// 3. It opens and closes connections to the actual underlying database, via the driver.
	// 4. It manages a pool of connections as needed
	//
	// It's safe for concurrent use by multiple goroutines. 协程安全
	// 原理是：DB拥有一个连接池，一旦有SQL操作，它将为该操作分配一个连接。
	// NOTE: 但是，DB并不清楚用户何时会结束使用，因为，如果返回多行数据时，database/sql要求流式获取
	//       即db.Query返回的是一个cursor(光标), 由用户自己一条条读取（当然底层可能不是一条条从Server端获取）
	// 类似：db.QueryRow 或者 db.Exec的请求，连接资源可以自动被回收的. 因为他们要么不返回，要么只返回一条数据
	//
	// Open may just validate its arguments without creating a connection to the database.
	// To verify that the data source name is valid, call Ping. 发送了真正的请求
	// DB(sql.Open()) is meant to be something that is essentially global
	// for the lift of your program. DB should be long-lived.
	//
	// If you don’t treat the sql.DB as a long-lived object, you could experience problems
	// such as poor reuse and sharing of connections, running out of available network resources,
	// or sporadic failures due to a lot of TCP connections remaining in TIME_WAIT status.
	// Such problems are signs that you’re not using database/sql as it was designed.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	if err != nil {
		panic(err.Error())
	}
	// NOTE: It is rare to Close a DB, as the DB handle is meant to be long-lived
	//       and shared between many goroutines.
	// defer db.Close()

	// NOTE: you should, in general, always prepare queries to be used multiple times
	// 1. the stmt is not bound to a single connection. 这是与db.Query的一个区别
	// 2. 在其它数据库中，statement的含义是：
	//	  首先指定一个带参数的命令给DB；DB返回一个编号；用户多次执行该statement时，只需要指定编号即可。
	//    所有这些操作在一个连接中完成
	// 3. database/sql中的statement，会依据底层的连接忙闲程度，适当时使用其它连接进行操作.
	stmtIns, err := db.Prepare("INSERT INTO user (age) VALUES( ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT age FROM user WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	// NOTE: db.Query, db.QueryRow ... 这些操作均适用于stmt
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}
