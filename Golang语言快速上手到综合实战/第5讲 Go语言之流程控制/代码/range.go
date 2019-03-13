// main.go
package main

import (
	"fmt"
)

func main() {
	x := "zhangsan"

	for _, v := range x {
		fmt.Printf("%c\n", v)
	}
}
