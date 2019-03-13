package main

import (
	"fmt"
)

func main() {
	// 声明函数类型
	type sum func(x, y int) int
	// 声明函数变量1
	/*
		func add(x, y int) int {
			return x + y
		}
		f := add
		fmt.Println(f(3,4))
	*/
	// 声明函数变量2,
	/*
		var f sum = func(x, y int) int {
			return x + y
		}
		fmt.Println(f(3, 4))
	*/
	// 声明函数变量3
	var result = func(x, y int) int {
		return x + y
	}(3, 4)
	fmt.Println(result)
}
