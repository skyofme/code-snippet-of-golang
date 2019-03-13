package main

import (
	"fmt"
)

//https://blog.go-zh.org/gos-declaration-syntax
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(3, 4))
}
