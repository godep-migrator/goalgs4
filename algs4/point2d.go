package algs4

type Point2D struct {
	X      float64
	Y      float64
	Canvas *Canvas
}

func NewPoint2D(x, y float64, canvas *Canvas) *Point2D {
	if canvas == nil {
		canvas = DefaultCanvas
	}

	return &Point2D{
		X:      x,
		Y:      y,
		Canvas: canvas,
	}
}

func (me *Point2D) Draw() {
	return
}
