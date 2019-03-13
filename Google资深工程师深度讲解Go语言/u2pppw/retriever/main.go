package main

import (
	"fmt"
	"time"

	"learngo/retriever/mock"
	"learngo/retriever/real"
)

// 由使用者决定要实现什么方法
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string ) string
}

const url = "http://www.imooc.com"
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string {
			"name":"ccmouse",
			"course":"golang",
		})
}

// 接口组合，类似于接口继承接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url,map[string]string {
		"contents" : "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(r)
	/*
	mock.Retriever {%!V(string=this is a fake imooc.com)}
	*real.Retriever &{%!V(string=Mozilla/5.0) %!V(time.Duration=60000000000)}
	*/

	//fmt.Println(download(r))

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if 	mockRetriever, ok := r.(*mock.Retriever); ok != false {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))

	// 使用interface{} 来代替类型声明以实现泛型
	var num int
	var v interface{}
	//fmt.Println(v.(int))

	fmt.Printf("%T %T", num, v)


}

func inspect(r Retriever) {
	fmt.Printf("%T %V\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Print("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Print("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
