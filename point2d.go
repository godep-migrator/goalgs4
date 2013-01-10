package algs4

type Point2D struct {
	X float64
	Y float64

	canvas *Canvas
}

func NewPoint2D(x, y float64, canvas *Canvas) *Point2D {
	if canvas == nil {
		canvas = DefaultCanvas
	}

	return &Point2D{
		X: x,
		Y: y,

		canvas: canvas,
	}
}

func (me *Point2D) Draw() {
	me.canvas.Pixel(me.X, me.Y)
}
