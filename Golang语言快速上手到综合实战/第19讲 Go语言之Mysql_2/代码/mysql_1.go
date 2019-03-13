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

	stmt, err := db.Prepare("update user_info set username=? where id=?")

	res, err := stmt.Exec("lisi", 22)
	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
