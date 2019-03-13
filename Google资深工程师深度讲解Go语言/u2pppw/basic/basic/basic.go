package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
使用var关键字
var a, b, c bool
var s1, s2 string = "hello", "world"
可放在函数内，或直接放在包内

使用编译器自动决定类型
var a, b, c = 1, true, "string"

使用:= 定义变量，只能在函数内使用
a := 3

内建变量类型

bool string

(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr（长度由操作系统决定）

byte, rune(类char，一字节)

float32, float64, complex64, complex128

*/

// 零值
func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)
}

// 初始化
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"

	// fmt.Printf("%d %q\n", a, s)
	fmt.Println(a, b, s)
}

// 推断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 作用域,不是全局变量！
var a = 6
var d = false
var (
	f = true
	g = "tank"
)
// 在函数外面不能使用：= 声明并赋值， 必须有var func 声明
// e := "def"
func variablePrint() {
	fmt.Println(a, d, f, g)
}

// 简短声明
func variableShorter() {
	// 作用域，作用在这个方法内
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func cmplxs() {
	c := 3 + 4i
	fmt.Print(c)
	fmt.Printf("= %v", cmplx.Abs(c))
}

// 欧拉公式
func eular() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f \n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

// 类型转换，go只有显式类型转换，没有隐式类型转换,小瑕疵：float-> int 丢失精度，解决：+ 0.5 向下取整？
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
// 其他语言常量大写，这里不大写，因为：go中的大写字母开头（代表着public）意味着其他包也能使用这些东西（变量，常量，函数等）
const filename = "abc.txt"

func consts() {

	const (
		a, b = 3, 4
		d    = 5
	)
	var c int
	// 里面不需要强转，因为const类型的数据被看作是文本
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, d, filename)
}

// 特殊的常量:枚举
func enums() {
	// 通过iota实现从0开始的自增长
	const (
		//cpp    = 0
		cpp    = iota
		java   = 1
		python = 3
		golang = 4
	)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello Golang")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variablePrint()
	variableShorter()
	cmplxs()
	eular()
	triangle()
	consts()
	enums()
}
