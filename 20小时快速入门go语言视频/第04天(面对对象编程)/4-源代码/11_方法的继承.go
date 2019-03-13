package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//Person指针类型，实现了一个方法
func (tmp *Person) PrintInfoPointer() {
	fmt.Printf("person pointer name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//Person值类型， 实现了一个方法
func (tmp Person) PrintInfoValue() {
	fmt.Printf("person value name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//Student指针类型，实现了一个方法
func (tmp *Student) PrintInfoPointer() {
	fmt.Printf("student pointer name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//Person值类型， 实现了一个方法
func (tmp Student) PrintInfoValue() {
	fmt.Printf("student value name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	// 如果没有重写，则继承并调用超类的方法；如果重写了超类的方法，则调用自己的方法
	s.PrintInfoPointer()
	s.PrintInfoValue()

	// 显式调用超类的方法
	s.Person.PrintInfoPointer()
	s.Person.PrintInfoValue()

}
