package algs4

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomUniform(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func Random() float64 {
	return rand.Float64()
}
