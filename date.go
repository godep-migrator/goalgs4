package algs4

import (
	"fmt"
)

type Date struct {
	Month int
	Day   int
	Year  int
}

func NewDate(m, d, y int) *Date {
	return &Date{
		Month: m,
		Day:   d,
		Year:  y,
	}
}

func (me *Date) String() string {
	return fmt.Sprintf("%d/%d/%d", me.Month, me.Day, me.Year)
}
