package main

import "fmt"

func split(sum int) (x, y int) {
	// 求个位
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
