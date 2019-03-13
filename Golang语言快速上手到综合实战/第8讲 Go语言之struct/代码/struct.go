// main.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	age        int
	speciality string
}

func main() {
	student := Student{Person{"zhangsan", 25}, 30, "maths"}
	fmt.Printf("%v", student.Person.age)
}
