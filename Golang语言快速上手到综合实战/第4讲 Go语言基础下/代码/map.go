// map
package main

import (
	"fmt"
)

func main() {
	/*
		var x map[string]float32
		x["zhangsan"] = 64.2
		fmt.Printf("%v", x)
		上述程序会报错 因为map使用前必须make分配内存 否则会报 nil map错误
	*/

	//正确用法如下

	//正常声明
	var x map[string]float32
	x = make(map[string]float32)
	x["zhangsan"] = 64.2
	fmt.Printf("%v\n", x)

	//或者直接使用简短声明符号
	y := make(map[string]float32)
	y["lisi"] = 76.5
	fmt.Printf("%v\n", y)
}
