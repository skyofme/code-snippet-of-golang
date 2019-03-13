package container

import "fmt"

func updateSlice(s []int) {
	// 数组也跟着改变
	s[0] = 100
}

func main() {
	// slice 本身是没有数据的，可以看作是数组的一个视图
	arr := [...]int{0,1,2,3,4,5,6,7}

	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

	s1 := arr[2:]
	fmt.Println(s1)
	// 传入切片对数组进行操作，可以省略*和&
	s2 := arr[:]
	fmt.Println(s2)

	fmt.Println("After updateSlice(s1)")

	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println(s2)
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	// reslice
	s2 = s2[:5]
	fmt.Println(s2)

	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	// slice可以向后扩展，不可以向后扩展
	s2 = s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
	// s[i] 不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	//fmt.Println(s2[5])
	//fmt.Println(s2[2:7])
	fmt.Println(len(arr))
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	fmt.Println("Appending slice")

	s3 := append(s2, 10)
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组，并复制原数组的值
	// 原数组如果不被使用将被垃圾回收机制回收
	// append后，ptr,len,cap 都可能改变
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println(arr)
}
