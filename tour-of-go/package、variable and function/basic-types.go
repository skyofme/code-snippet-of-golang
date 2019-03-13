package main

import (
	"fmt"
	"math/cmplx"
)

/*
bool

string

 int int8  int16  int32  int64
uint uint8 uint16 uint32 uint64

byte // uint8 别名

rune // int32 别名，表示一个Unicode码点

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("%T %v\n", ToBe, ToBe)
	fmt.Printf("%T %v\n", MaxInt, MaxInt)
	fmt.Printf("%T %v\n", z, z)
}
