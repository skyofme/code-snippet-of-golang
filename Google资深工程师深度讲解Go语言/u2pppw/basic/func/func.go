package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 忽略某个返回值
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op,
		)
		// panic("unsupported operation:" + op)
	}
}

// 13 / 3 = 4 ... 1
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数式编程 改写 eval
func apply(op func(int, int) int, a, b int) int {
	// 获得函数的指针
	p := reflect.ValueOf(op).Pointer()
	// 获取函数名：包名.函数名.匿名函数名(funcN)
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

// 重写math.Pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(numbers ...int) (sum int) {
	for i := range numbers {
		sum += numbers[i]
	}
	return
}

// 探讨值传递：传递值和传递指针
func swap0(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	// 处理error
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "*"))

	// 带余数除法
	q, r := div(3, 2)
	fmt.Println(q, r)

	// 函数式编程调用1: 定义并调用函数
	result := apply(pow, 3, 4)
	fmt.Println(result)

	// 函数式编程调用2: 使用匿名函数，使函数不在包级别
	result = apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 5,
	)
	fmt.Println(result)

	// 带可变参数的函数的调用
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 值传递的探讨
	a, b := 3, 4
	swap0(&a, &b)
	fmt.Println(a, b)

	c, d := swap(a, b)
	fmt.Println(c, d)

}
