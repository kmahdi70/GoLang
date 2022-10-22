package main

import (
	"fmt"
)

type Student struct {
	fn  string
	ln  string
	age int
}

func main() {
	var st []Student

	/* student[0].fn = "ali"
	student[0].ln = "rezaei"
	student[0].age = 34
	*/
	st = append(st, Student{fn: "alireza", ln: "ahmadi", age: 34})
	st = append(st, Student{fn: "alireza", ln: "ahmadi", age: 34})
	st = append(st, Student{fn: "alireza", ln: "ahmadi", age: 34})
	st[0].fn = "ali"
	fmt.Println(st)

}
