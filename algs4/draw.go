package algs4

import (
	"errors"
	"image"
	"image/draw"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
)

func Draw() error {
	img := DefaultCanvas.Image()
	if img == nil {
		return errors.New("DefaultCanvas's image is nil!")
	}

	w, err := x11.NewWindow()
	if err != nil {
		return err
	}

	draw.Draw(w.Screen(), w.Screen().Bounds(), img, image.ZP, draw.Src)
	w.FlushImage()

	for evt := range w.EventChan() {
		switch evt := evt.(type) {
		case ui.KeyEvent:
			if evt.Key == ' ' {
				return nil
			}
		}
	}
	return nil
}
