// select.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Print("come to read ch!")
	case <-time.After(time.Second):
		fmt.Print("come to timeout!")
	}

	fmt.Print("end of code!")
}
