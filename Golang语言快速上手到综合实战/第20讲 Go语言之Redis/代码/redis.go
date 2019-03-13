// redis
package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("hello beifeng"))
	if err != nil {
		panic(err)
	}

	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	//test hmset
	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "male"

	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}

	//test zset
	_, err = client.Zadd("test_zset", []byte("beifeng"), 100)
	if err != nil {
		panic(err)
	}
}
