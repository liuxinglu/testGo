package main

import (
	"fmt"
	"net/http"
)

func insertStudent(w http.ResponseWriter, r *http.Request) {
	accessControl(w)
	r.ParseForm() //解析参数，默认是不会解析的
	id := insertUser(r.Form["name"][0], r.Form["age"][0], r.Form["sex"][0])
	s := Stud{id, r.Form["name"][0], r.Form["age"][0], r.Form["sex"][0]}
	student := Student{"insertStudent", s}
	fmt.Fprintf(w, "{\"%s\":\"%s\", \"%s\":\"%d\", \"%s\":\"%s\", \"%s\":\"%s\", \"%s\":\"%s\"}",
		"code", student.Code,
		"stuId", student.s.stuId, "name", student.s.name, "age", student.s.age, "sex", student.s.sex)
}

func findStudent(w http.ResponseWriter, r *http.Request) {
	accessControl(w)
	r.ParseForm()
	users := findUser(r.Form["name"][0])
	s := Stud{users[0].stuId, users[0].name, users[0].age, users[0].sex}
	student := Student{"findStudent", s}
	fmt.Fprintf(w, "{\"%s\":\"%s\", \"%s\":\"%d\", \"%s\":\"%s\", \"%s\":\"%s\", \"%s\":\"%s\"}",
		"code", student.Code,
		"stuId", student.s.stuId, "name", student.s.name, "age", student.s.age, "sex", student.s.sex)
}

func accessControl(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
}
