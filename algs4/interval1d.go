package algs4

type Interval1D struct {
	Lo float64
	Hi float64
}

func NewInterval1D(lo, hi float64) *Interval1D {
	return &Interval1D{
		Lo: lo,
		Hi: hi,
	}
}
