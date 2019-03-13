// mysql.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT user_info  SET username=?,departname=?,create_time=?")

	res, err := stmt.Exec("zhangsan", "技术部", "2016-12-09")
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
