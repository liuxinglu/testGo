package main

type Stud struct {
	stuId int64
	name  string
	age   string
	sex   string
}

type Student struct {
	Code string
	s    Stud
}
