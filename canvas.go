package algs4

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"math"
	"os"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
)

var (
	DefaultCanvas, DefaultCanvasError = NewCanvas(512, 512,
		float64(0.0), float64(1.0), float64(0.0), float64(1.0))

	halfWidthNegativeError  = errors.New("half width can't be negative")
	halfHeightNegativeError = errors.New("half height can't be negative")
)

func init() {
	if DefaultCanvasError != nil {
		fmt.Fprintf(os.Stderr,
			"Failed to initialize DefaultCanvas: %v\n", DefaultCanvasError)
		panic(DefaultCanvasError)
	}
}

type Canvas struct {
	Height int
	Width  int
	Xmin   float64
	Xmax   float64
	Ymin   float64
	Ymax   float64

	img    *image.RGBA
	window ui.Window
}

func NewCanvas(height, width int, xmin, xmax, ymin, ymax float64) (*Canvas, error) {
	canvas := &Canvas{
		Height: height,
		Width:  width,
		Xmin:   xmin,
		Xmax:   xmax,
		Ymin:   ymin,
		Ymax:   ymax,

		img:    image.NewRGBA(image.Rect(0, 0, height, width)),
		window: nil,
	}

	return canvas, nil
}

func (me *Canvas) ensureHasWindow() error {
	if me.window != nil {
		return nil
	}

	window, err := x11.NewWindow()
	if err != nil {
		return err
	}

	me.window = window

	return nil
}

func (me *Canvas) Image() *image.RGBA {
	return me.img
}

func (me *Canvas) Bounds() image.Rectangle {
	return me.img.Bounds()
}

func (me *Canvas) Rectangle(x, y, halfWidth, halfHeight float64) error {
	if halfWidth < 0 {
		return halfWidthNegativeError
	}

	if halfHeight < 0 {
		return halfHeightNegativeError
	}

	f2 := float64(2.0)

	xs := me.scaleX(x)
	ys := me.scaleY(y)
	ws := me.factorX(f2 * halfWidth)
	hs := me.factorY(f2 * halfHeight)

	if ws <= 1 && hs <= 1 {
		me.Pixel(x, y)
		return nil
	}

	originX, originY := int(xs-ws/f2), int(ys-hs/f2)
	draw.Draw(me.img,
		image.Rect(originX, originY, originX+int(ws), originY+int(hs)),
		&image.Uniform{image.White}, image.ZP, draw.Src)

	return nil
}

func (me *Canvas) Pixel(x, y float64) error {
	err := me.ensureHasWindow()
	if err != nil {
		return err
	}

	originX, originY := int(me.scaleX(x)), int(me.scaleY(y))
	draw.Draw(me.img, image.Rect(originX, originY, originX+1, originY+1),
		&image.Uniform{image.White}, image.ZP, draw.Src)
	me.window.FlushImage()

	return nil
}

func (me *Canvas) Draw() error {
	err := me.ensureHasWindow()
	if err != nil {
		return err
	}

	draw.Draw(me.window.Screen(), me.window.Screen().Bounds(), me.img, image.ZP, draw.Src)
	me.window.FlushImage()

	for evt := range me.window.EventChan() {
		switch evt := evt.(type) {
		case ui.KeyEvent:
			if evt.Key == 'q' {
				return nil
			}
		}
	}
	return nil
}

func (me *Canvas) scaleX(x float64) float64 {
	return float64(me.Width) * (x - me.Xmin) / (me.Xmax - me.Xmin)
}

func (me *Canvas) scaleY(y float64) float64 {
	return float64(me.Height) * (me.Ymax - y) / (me.Ymax - me.Ymin)
}

func (me *Canvas) factorX(w float64) float64 {
	return w * float64(me.Width) / math.Abs(me.Xmax-me.Xmin)
}

func (me *Canvas) factorY(h float64) float64 {
	return h * float64(me.Height) / math.Abs(me.Ymax-me.Ymin)
}
