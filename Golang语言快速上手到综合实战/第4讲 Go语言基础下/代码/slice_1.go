// slice_1
package main

import (
	"fmt"
)

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var y []int
	y = x[1:3]
	fmt.Printf("%v\n", y)

	y = x[:3] //左边不写 代表从0开始
	fmt.Printf("%v\n", y)

	y = x[7:] //右边不写代表一直到最后
	fmt.Printf("%v\n", y)
}
