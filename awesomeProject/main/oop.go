package main

import (
    "fmt"
)

/* func compare(a,b int) bool {
    return a < b
}*/

type Point struct {
    px float32
    py float32
}

// 接收者， 参数， 无返回值
func (point *Point) setXY(px, py float32) {
    point.px = Px
    point.py = py
}

// 接收者， 无参数， 返回值
func (point *Point) getXY() (float32, float32) {
    return point.px, point.py
}

type Integer struct {
    value int
}

type (a Integer) compare(b Integer) bool {
    return a.value < b.value
}

func main() {
    point := new(Point)
    point.setXY(1.23, 4.56)
    px, py:= point.getXY()
    fmt.Print(px,py)
}
