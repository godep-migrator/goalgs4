package main

import (
	"fmt"
	"time"

	. "goalgs4"
)

func main() {
	for n := int64(250); true; n += n {
		cnt, dur := timeTrial(n)
		fmt.Printf("%7v %5v (%v)\n", n, cnt, dur)
	}
}

func timeTrial(n int64) (int64, time.Duration) {
	max := float64(1000000)
	a := []int64{}

	for i := int64(0); i < n; i++ {
		a = append(a, int64(RandomUniform(-max, max)))
	}

	timer := NewStopwatch()
	cnt := ThreeSumCount(a)
	return cnt, timer.ElapsedTime()
}
