package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, hight int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{img.width, img.hight}}
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{100, 200}
	pic.ShowImage(m)
}
