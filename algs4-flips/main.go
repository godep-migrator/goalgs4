package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

import (
	"github.com/meatballhat/box-o-sand/algs4/src/go/algs4"
)

const USAGE string = "Usage: algs4-flips <nflips>"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	nflips, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("Wat.", err)
		os.Exit(2)
	}

	heads := algs4.NewCounter("heads")
	tails := algs4.NewCounter("tails")

	randMax := big.NewInt(2)

	for t := int64(0); t < nflips; t += int64(1) {
		n, err := rand.Int(rand.Reader, randMax)
		if err != nil {
			fmt.Println("Ugh:", err)
			os.Exit(4)
		}

		if n.Int64() < 1 {
			heads.Increment()
		} else {
			tails.Increment()
		}
	}

	fmt.Println(heads)
	fmt.Println(tails)

	d := heads.Tally() - tails.Tally()

	fmt.Printf("delta: %d\n", int(math.Abs(float64(d))))
}
