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

	rows, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var username string
		var department string
		var create_time string
		err = rows.Scan(&id, &username, &department, &create_time)
		fmt.Println(id, username, department, create_time)
	}
}
