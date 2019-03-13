package main

import (
	"fmt"
	"strconv"
	"time"
)

// 参数列表：两个整型，一个传递整型的通道
func Add(x, y int, quit chan int) {
	z := x + y
	fmt.Println(z)
	quit <- 1
}

func Read(ch chan int) {
	// 接收通道ch所传递的值
	value := <-ch
	// 迭代打印value的值
	fmt.Println("value : " + strconv.Itoa(value))
}

func Write(ch chan int) {
	// 让ch传递10这个值
	ch <- 10
}

func main() {
	// 创建通道,没有限定通道值的长度，所以是非缓冲
	ch := make(chan int)
	// 并发执行读写函数
	go Read(ch)
	go Write(ch)

	time.Sleep(10)

	fmt.Println("end of code.")

	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Add(i, i, chs[i])
	}

	for _, v := range chs {
		<-v
	}
}
