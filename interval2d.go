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
	return me.Xint.Contains(p.X) && me.Yint.Contains(p.Y)
}

func (me *Interval2D) Draw() {
	xc := (me.Xint.Left + me.Xint.Right) / 2.0
	yc := (me.Yint.Left + me.Yint.Right) / 2.0
	me.Canvas.Rectangle(xc, yc, me.Xint.Length()/2.0, me.Yint.Length()/2.0)
}

func (me *Interval2D) Area() float64 {
	return me.Xint.Length() * me.Yint.Length()
}
