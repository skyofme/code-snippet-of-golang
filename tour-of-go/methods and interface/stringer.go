package main

import "fmt"

/*
一个普遍存在的接口是 fmt 包中定义的 Stringer。

type Stringer struct {
    String() string
}
Stringer 是一个可以用字符串描述自己的类型。
`fmt`包 （还有许多其他包）使用这个来进行输出。
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)
}
