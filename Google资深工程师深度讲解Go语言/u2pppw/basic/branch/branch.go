package basic

import (
	"fmt"
	"io/ioutil"
)

// 参数类型写在参数名后面，函数返回值也是类似
func grade(score int) (g string) {
	// switch 分支语法，fallthrough,直接将控制权给下一个case，该关键字必须位于当前case（不为最后一个case）的最后一句
	switch {
	case score <0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return
}

func main() {
	const filename = "abc.txt"
	/*
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(contents)
		fmt.Printf("%s \n", contents)
	}
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(contents)
		fmt.Printf("%s \n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(61),
		grade(79),
		grade(81),
		grade(89),
		grade(99),
		grade(100),
		//grade(101),
	)
}
