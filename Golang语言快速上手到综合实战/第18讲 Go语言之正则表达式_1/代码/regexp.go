// regexp
package main

import (
	"fmt"
	"regexp"
)

func main() {
	//reg := regexp.MustCompile("\\w+") 正则表达式中的\需要转义
	reg := regexp.MustCompile(`^z.*l$`)

	result := reg.FindAllString("zhangsanl", -1)
	fmt.Printf("%v\n", result)

	reg1 := regexp.MustCompile(`^z(.*)l$`)

	result1 := reg1.FindAllString("zhangsand", -1)
	fmt.Printf("%v\n", result1)

	reg2 := regexp.MustCompile(`^z(.{1})(.{1})(.*)l$`)

	result2 := reg2.FindAllStringSubmatch("zhangsanl", -1)
	fmt.Printf("%v\n", result2)
}
