package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (myImage Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (image Image) At(x int, y int) color.Color {
	v := uint8((x + y) / 2)
	// v := uint8(x * y)
	// v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
