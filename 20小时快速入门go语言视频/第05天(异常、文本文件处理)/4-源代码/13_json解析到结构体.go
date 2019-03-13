package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {

	jsonBuf := `
	{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
}`

	var tmp IT //定义一个结构体变量
	// 将json转换成string-{}interface并存储在结构体中
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"` //二次编码
	}

	var tmp2 IT2
	// 若结构体中不存在key所对应的字段，则舍去
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)

	type IT3 struct {
		Company  string   `json:"company"`
		Subjects []string `json:"subjects"` //二次编码
		IsOk     bool     `json:"isok"`
		Price    float64  `json:"price"`
		Other    string   `json:"exist?"`
	}

	var temp3 IT3
	// 若json串中不存在结构体相应字段，则相应字段赋值为零值，key用结构体的字段名
	err = json.Unmarshal([]byte(jsonBuf), &temp3)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("temp3 = %+v\n", temp3)
}
