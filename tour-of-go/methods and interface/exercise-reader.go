package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// TODO: Add a Read(byte) (int, error) method to MyReader.
// 实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
func main() {
	reader.Validate(MyReader{})
}