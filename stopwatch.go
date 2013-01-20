package algs4

import (
	"time"
)

type Stopwatch struct {
	start time.Time
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		start: time.Now(),
	}
}

func (me *Stopwatch) ElapsedTime() time.Duration {
	return time.Now().Sub(me.start)
}
