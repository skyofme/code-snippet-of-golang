package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Printf("%d ", n)
		}(func() int {
			n := fibonacci(i)
			fmt.Printf("%d ", n)
			return n
		}())
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
