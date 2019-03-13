// panic.go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("After panic!") //panic之后 defer里面的依然可以得到执行
	}()

	panic("I am panicing!")

	fmt.Println("After panic!") //panic后面得不到执行
}
