package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 实现stringer接口，重写string方法，类似toString()
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = "Post: " + form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return "Get: " + r.Contents
}
