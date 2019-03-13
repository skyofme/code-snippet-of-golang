package container

import "fmt"

// 数组是值类型:长度不一致的数组不能作为参数传递(拷贝数组)，在函数中的修改不能传递到原数组中
func printArray(arr [5]int) {
	for i,v  := range arr {
		fmt.Println(i,v)
	}
}

func changeArray(arr *[5]int) {
	for i,v  := range arr {
		fmt.Println(i,v)
	}
}

func main() {
	// 声明一个数组
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 二维数组，四行五列
	var grid [4][5]bool

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 遍历下标
	for i := range arr3 {
		fmt.Println(i)
	}

	sum := 0
	// 遍历元素
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	//maxIndex, maxValue := -1, -1
	maxIndex := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxIndex, maxValue = i, v
		}
	}
	fmt.Println(maxIndex, maxValue)

	printArray(arr1)
	printArray(arr3)

	changeArray(&arr1)
	changeArray(&arr3)
}
