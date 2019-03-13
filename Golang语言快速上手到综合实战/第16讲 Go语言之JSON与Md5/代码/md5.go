// json.go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("zhangsan"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
