package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	. "github.com/meatballhat/goalgs4"
)

const USAGE string = "Usage: algs4-flips <nflips>"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	nflips, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Wat.", err)
		os.Exit(2)
	}

	heads := NewCounter("heads")
	tails := NewCounter("tails")

	for t := 0; t < nflips; t++ {
		if Random() > 0.5 {
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
