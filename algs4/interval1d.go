package algs4

type Interval1D struct {
	Left  float64
	Right float64
}

func NewInterval1D(left, right float64) *Interval1D {
	return &Interval1D{
		Left:  left,
		Right: right,
	}
}

func (me *Interval1D) Contains(x float64) bool {
	return me.Left <= x && x <= me.Right
}

func (me *Interval1D) Length() float64 {
	return me.Right - me.Left
}
