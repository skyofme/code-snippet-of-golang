package main

import "fmt"

type Node struct {
	value       int
	Left, Right *Node
}

// 函数定义
func Print(node Node) {
	fmt.Print(node)
}

// 方法定义
func (node Node) Print() {
	fmt.Print(node)
}

// 工厂函数，自定义构造函数,返回局部变量的地址（闭包？溢出？）
func CreateNode(value int) *Node {
	// 编译器和运行环境决定局部变量是分配在堆中还是栈中，我们没必要知道，因为有垃圾回收机制
	// 给外面用了，则可能在堆中分配，反之可能在栈中分配
	return &Node{value:value}
}

// 指针接收者可以调用值接收者和本身的方法，值接收者只能调用本身的方法
// 改变必须使用指针，结构体过大也用指针，如果要用，建议所有方法都用指针接收者
// 只有使用指针接收者才能改变结构体，nil也可以调用指针接收者的方法
func (node *Node) SetValue(value int) {
	node.value = value
}

// 中序遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 组合实现后序遍历
type myNode struct {
	node *Node
}

// 后序遍历
func (myNode *myNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myNode{myNode.node.Left}
	right := myNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}


func main() {
	var root Node

	root = Node{value:3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Right.Right = CreateNode(2)

	nodes := []Node {
		{value:3},
		{},
		{6,nil,nil},
	}
	fmt.Println(nodes)

	// 调用方法
	root.Print()
	// 调用函数
	Print(root)

	root.Right.SetValue(4)
	root.Right.Print()

	root.Traverse()
}


