package main

import (
	"fmt"
)

// fibonacci 函数会返回一个返回int的函数
/*
实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
*/

func fibonacci() func() int {

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
