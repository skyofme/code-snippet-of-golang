package main

import (
	"fmt"
)

func main() {
	ch2 := make(chan string, 1)
	// 启用一个Goroutine来并发执行代码块的方法
	// 关键字go 后面跟需要并发执行的代码块，一般由一个匿名函数代表
	go func() {
		ch2 <- ("already send.")
	}()
	var value string = "data : "
	value = value + (<-ch2)
	fmt.Println(value)
}
