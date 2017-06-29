package main

import (
	"log"
	"net/http"
)

func main() {
	db = connectSql("root", "liuxinglu", "test")
	http.HandleFunc("/insertStudent/", insertStudent) //插入学生
	http.HandleFunc("/findStudent/", findStudent)     //查找学生
	err := http.ListenAndServe(":9090", nil)          //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
