package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		// input
		timeout <- 1
	}()

	select {
	// recieve
	case <-ch:
		fmt.Print("come to read ch ")
	case <-timeout:
		fmt.Print("come to timeout ")
	}
	fmt.Print("end of code")
}