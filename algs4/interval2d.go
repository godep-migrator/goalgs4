package algs4

type Interval2D struct {
	Xint   *Interval1D
	Yint   *Interval1D
	Canvas *Canvas
}

func NewInterval2D(xint, yint *Interval1D, canvas *Canvas) *Interval2D {
	if canvas == nil {
		canvas = DefaultCanvas
	}

	return &Interval2D{
		Xint:   xint,
		Yint:   yint,
		Canvas: canvas,
	}
}

func (me *Interval2D) Contains(p *Point2D) bool {
	return false
}

func (me *Interval2D) Draw() {
	return
}

func (me *Interval2D) Area() float64 {
	return float64(0)
}
