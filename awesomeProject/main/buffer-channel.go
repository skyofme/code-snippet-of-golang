package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {
	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println("come to end goroutine 1")
}

func main() {
	// make(chan int, 0) == make(chan int) 都是非缓冲的channel
	ch = make(chan int, 0)

	// make(chan int, 2) 是带缓冲的channel
	ch = make(chan int, 2)

	go test_channel()

	time.Sleep(2 * time.Second)

	// main函数结束后不执行其他的channel！！！
	fmt.Println("running end!")

	time.Sleep(time.Second)
}
