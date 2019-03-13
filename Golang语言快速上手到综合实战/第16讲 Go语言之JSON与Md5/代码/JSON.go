// md5.go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"student_name"`
	Age  int
}

func main() {
	//对数组类型的json编码
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))

	//对map类型进行json编码
	m := make(map[string]float64)
	m["zhangsan"] = 100.4
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(s2))

	//对对象进行json编码
	student := Student{"zhangsan", 26}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println(string(s3))

	//对s3进行json解码
	var s4 interface{}
	json.Unmarshal([]byte(s3), &s4)
	fmt.Printf("%v", s4)
}
