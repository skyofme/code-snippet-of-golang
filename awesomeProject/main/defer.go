package main

import (
	"fmt"
)

//注意，当一个函数中存在多个defer语句时，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序。下面的示例可以很好的证明这一点：
func deferIt() {
	defer func() {
		fmt.Print(1)
	}()
	defer func() {
		fmt.Print(2)
	}()
	defer func() {
		fmt.Print(3)
	}()
	fmt.Print(4)
}

//deferIt函数的执行会使标准输出上打印出4321。请大家猜测下面这个函数被执行时向标准输出打印的内容，并真正执行它以验证自己的猜测。

//最后论证一下自己的猜测为什么是对或者错的。
func deferIt2() {
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/**
1. defer携带的表达式语句代表的是对某个函数或方法的调用。
这个调用可能会有参数传入，比如：fmt.Print(i + 1)。
如果代表传入参数的是一个表达式，那么在defer语句被执行的时候该表达式就会被求值了。注意，这与被携带的表达式语句的执行时机是不同的。请揣测下面这段代码的执行：
*/
func deferIt3() {
	f := func(i int) int {
		fmt.Printf(" %d ", i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}

//它在被执行之后，标准输出上打印出1 2 3 4 40 30 20 10 。

/**
2. 如果defer携带的表达式语句代表的是对匿名函数的调用，那么我们就一定要非常警惕。请看下面的示例：
*/
/*
func deferIt4() {
	for i := 1; i < 5; i++ {
		defer func() {
			fmt.Print(i)
		}()
	}
}
*/
/**
deferIt4函数在被执行之后标出输出上会出现5555，而不是4321。
原因是defer语句携带的表达式语句中的那个匿名函数包含了对外部（确切地说，是该defer语句之外）的变量的使用。
注意，等到这个匿名函数要被执行（且会被执行4次）的时候，包含该defer语句的那条for语句已经执行完毕了。
此时的变量i的值已经变为了5。
因此该匿名函数中的打印函数只会打印出5。
正确的用法是：把要使用的外部变量作为参数传入到匿名函数中。

修正后的deferIt4函数如下：
*/
func deferIt4() {
	for i := 1; i < 5; i++ {
		defer func(n int) {
			fmt.Print(n)
		}(i)
	}
}

func main() {

	deferIt()

	deferIt2()

	deferIt3()

	deferIt4()
}
