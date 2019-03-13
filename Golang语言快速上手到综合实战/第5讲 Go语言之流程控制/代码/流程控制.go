// main
package main

import (
	"fmt"
)

func main() {
	x := 2
	switch x {
	case 1:
		fmt.Print("beifeng 1")
	case 2:
		fallthrough
	case 3:
		fmt.Print("beifeng 2")
	default:
		fmt.Print("beifeng 3")
	}
}
