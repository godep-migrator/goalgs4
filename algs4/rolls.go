package algs4

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const ROLL_SIDES = 6

func Rolls(nrolls int64) ([]*Counter, error) {
	var err error

	rolls := make([]*Counter, ROLL_SIDES+1)
	for i := 1; i <= ROLL_SIDES; i += 1 {
		rolls[i] = NewCounter(fmt.Sprintf("%d's", i))
	}

	var result *big.Int
	randMax := big.NewInt(ROLL_SIDES)

	for t := int64(0); t < nrolls; t += 1 {
		result, err = rand.Int(rand.Reader, randMax)

		if err != nil {
			return nil, err
		}

		rolls[int(result.Int64())+1].Increment()
	}

	return rolls, nil
}
