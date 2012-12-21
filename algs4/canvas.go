package algs4

import (
	"image"
)

var (
	DefaultCanvas = NewCanvas(512, 512)
)

type Canvas struct {
	Height int
	Width  int
	img    *image.RGBA
}

func NewCanvas(height, width int) *Canvas {
	return &Canvas{
		Height: height,
		Width:  width,
		img:    image.NewRGBA(image.Rect(0, 0, height, width)),
	}
}

func (me *Canvas) Image() *image.RGBA {
	return me.img
}

func (me *Canvas) Rectangle(x, y, halfWidth, halfHeight float64) {
	return
}
