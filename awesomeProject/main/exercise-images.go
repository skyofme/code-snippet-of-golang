package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
)

/*

还记得之前编写的图片生成器吗？现在来另外编写一个，不过这次将会返回 image.Image 来代替 slice 的数据。

自定义的 Image 类型，要实现必要的方法，并且调用 `pic.ShowImage`。

Bounds 应当返回一个 `image.Rectangle`，例如 `image.Rect(0, 0, w, h)`。

ColorModel 应当返回 `color.RGBAModel`。

At 应当返回一个颜色；在这个例子里，在最后一个图片生成器的值 v 匹配 `color.RGBA{v, v, 255, 255}`。

*/


type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
