package main

import (
	"fmt"
	"math/rand"
)

// 伪随机数，需要为随机数生成器指定不同的种子，比如时间
func main() {
	fmt.Println("random number is",
		rand.Intn(10))
}
