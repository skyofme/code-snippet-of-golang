package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 需求：将十进制数转换为二进制数

// for 循环讲解
func convertToBin(n int) (result string) {
	if (n == 0) {
		return ""
	}
	for ; n > 0; n /= 2 {
		// 取最低位
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

// 类while循环讲解
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 逐行读取
	scanner := bufio.NewScanner(file)

	// while 循环
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("abv")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("abc.txt")

	forever()
}