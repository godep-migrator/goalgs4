package algs4

import (
	"fmt"
)

const ROLL_SIDES = 6

func Rolls(nrolls int) []*Counter {
	var rolls []*Counter

	for i := 0; i < ROLL_SIDES; i++ {
		rolls = append(rolls, NewCounter(fmt.Sprintf("%d's", i+1)))
	}

	randMax := float64(ROLL_SIDES)
	zero := float64(0.0)

	for t := 0; t < nrolls; t++ {
		result := int(RandomUniform(zero, randMax))
		rolls[result].Increment()
	}

	return rolls
}
