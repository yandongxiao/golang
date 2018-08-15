package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getRows() *sql.Rows {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	checkError(err)
	fmt.Println(db.Ping())

	stmt, err := db.Prepare("select * from user where age > ?")
	checkError(err)

	rows, err := stmt.Query(100)
	checkError(err)
	return rows
}

func main() {
	scan()
}

func scan() {
	rows := getRows()
	for rows.Next() {
		// If an argument has type *interface{}, Scan copies the value
		// provided by the underlying driver without conversion.
		// 例如，下面的例子中，底层driver返回的类型已经很好用了int64, []uint8, int64
		// When scanning from a source value of type []byte to *interface{}, a copy of the
		// slice is made and the caller owns the result. 不共享数据
		var a, b, c interface{}
		rows.Scan(&a, &b, &c)
		fmt.Printf("%T, %T, %T\n", a, b, c)
		fmt.Printf("id=%d, name=%s, age=%d", a, b, c)
	}
	fmt.Println(rows.Err())
}

func columnTypes() {
	rows := getRows()
	columTypes, err := rows.ColumnTypes()
	checkError(err)

	// sql的一个字段：`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键')
	for _, ctype := range columTypes {
		fmt.Printf("name:%s, ", ctype.Name())

		// returns the database system name of the column type.
		// If an empty string is returned the driver type name is not supported.
		fmt.Printf("type:%s, ", ctype.DatabaseTypeName())

		if strings.ToLower(ctype.DatabaseTypeName()) == "varchar" {
			// Length returns the column type length for variable length column types
			// such as text and binary field types.
			// If the column type is not variable length, such as an int, ok is false
			if len, ok := ctype.Length(); ok {
				fmt.Printf("length: %d", len)
			}
		}

		// Nullable returns whether the column may be null.
		// If a driver does not support this property ok will be false
		nullable, ok := ctype.Nullable()
		if ok {
			fmt.Printf("nullable:%v, ", nullable)
		}

		// fmt.Println(ctype.DecimalSize()) // 0 0 false. mysql驱动也不支持该字段

		// ScanType returns a Go type suitable for scanning into using Rows.Scan.
		// If a driver does not support this property ScanType will return
		// the type of an empty interface.
		fmt.Printf("reflect: %v\n", ctype.ScanType())
	}

	fmt.Println(rows.Columns())
}
