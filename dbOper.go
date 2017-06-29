package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectSql(user string, psw string, dbname string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+psw+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("连接成功")
	}
	return db
}

func insertUser(name string, age string, sex string) int64 {
	stms, _ := db.Prepare("insert user (user_name,user_age,user_sex) values (?,?,?)")
	res, err := stms.Exec(name, age, sex)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("插入成功")
	}
	id, _ := res.LastInsertId()
	return id
}

func findUser(name string) []Stud {
	rows, err := db.Query("select * from user where user_name = ?", name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("查询成功")
	}
	arr := []Stud{}
	for rows.Next() {
		var s Stud
		if err := rows.Scan(&s.stuId, &s.name, &s.age, &s.sex); err != nil {
			fmt.Println(err)
		}

		arr = append(arr, s)
	}
	return arr
}
