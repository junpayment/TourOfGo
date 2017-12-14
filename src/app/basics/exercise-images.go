package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct {}

func (p *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Image) Bounds() image.Rectangle {
	return image.Rect(1, 2, 3, 4)
}

func (p *Image) At(x, y int) color.Color {
	return color.RGBA{1, 2, 3, 4}
}

func main() {
	m := Image{}
	pic.ShowImage(&m)
}
