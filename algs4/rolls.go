package algs4

import (
	"fmt"
)

const ROLL_SIDES = 6

func Rolls(nrolls int64) ([]*Counter, error) {
	rolls := make([]*Counter, ROLL_SIDES+1)
	for i := 1; i <= ROLL_SIDES; i += 1 {
		rolls[i] = NewCounter(fmt.Sprintf("%d's", i))
	}

	randMax := float64(ROLL_SIDES)

	for t := int64(0); t < nrolls; t++ {
		rolls[int(RandomUniform(float64(0.0), randMax))+1].Increment()
	}

	return rolls, nil
}
