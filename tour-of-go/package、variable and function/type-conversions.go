package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var c float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(c)
	fmt.Println(x, y, z)

	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Println(i, f, u)
}
