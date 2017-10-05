package main

import (
	"database/sql"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "leogsouza:@/c9?charset=UTF8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("leogsouza", "Project Development", "2016-10-03")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Last inserted ID", id)

	// update
	stmt, err = db.Prepare("update userinfo set username= ? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("legosuz", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Rows affected", affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println("ID:", uid)
		fmt.Println("Username:", username)
		fmt.Println("Department:", department)
		fmt.Println("Created:", created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("Rows affected on delete", affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
