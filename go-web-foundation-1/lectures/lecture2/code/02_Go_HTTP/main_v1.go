package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHello1)
	http.HandleFunc("/bye", sayBye1)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 1.")
}

func sayBye1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye, this is version 1.")
}
