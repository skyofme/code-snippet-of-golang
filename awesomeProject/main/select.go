package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		// 未被初始化的channel不能使其case生效
		case e, ok := <-ch:
			if !ok {
				fmt.Println("end.")
				return
			}
			fmt.Println(e)
			close(ch)
		default:
			fmt.Println("no data.")
			ch <- 1
		}
	}
}
