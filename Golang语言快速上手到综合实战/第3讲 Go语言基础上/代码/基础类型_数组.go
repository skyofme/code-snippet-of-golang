// 基础类型_数组
package main

import (
	"fmt"
)

func main() {
	/*var x [10]int //直接声明
	x[0] = 1
	x[1] = 2*/

	x := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", x) //注意%v的运用 %v可以打印数组 map 字符串等各类值
}
