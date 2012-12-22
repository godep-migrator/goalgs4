package algs4

import "fmt"

type Accumulator struct {
	total float64
	n     int
}

func NewAccumulator() *Accumulator {
	return &Accumulator{
		float64(0.0),
		0,
	}
}

func (me *Accumulator) AddDataValue(val float64) {
	me.n++
	me.total += val
}

func (me *Accumulator) Mean() float64 {
	return me.total / float64(me.n)
}

func (me *Accumulator) String() string {
	return fmt.Sprintf("Mean (%d values): %7.5f", me.n, me.Mean())
}
