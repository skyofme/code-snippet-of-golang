// main
package main

import (
	"fmt"
)

/*func compare(a, b int) bool {
	return a < b
}*/

type Point struct {
	px float32
	py float32
}

func (point *Point) setXY(px, py float32) {
	point.px = px
	point.py = py
}

func (point *Point) getXY() (float32, float32) {
	return point.px, point.py
}

type Integer struct {
	value int
}

func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

func main() {
	point := new(Point)
	point.setXY(1.23, 4.56)
	px, py := point.getXY()
	fmt.Print(px, py)
}
