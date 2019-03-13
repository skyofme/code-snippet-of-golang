package main

import (
	"fmt"
	"net/http"

	"github.com/gpmgo/gopm/modules/log"
)

type Hello struct {

}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err.Error())
	}
}
