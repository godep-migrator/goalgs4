package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

import (
	"github.com/meatballhat/box-o-sand/algs4/src/go/algs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 6 {
		die(errors.New("We need 4 floats and an int!"))
	}

	rand.Seed(time.Now().UTC().UnixNano())

	xlo, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		die(err)
	}

	xhi, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		die(err)
	}

	ylo, err := strconv.ParseFloat(os.Args[3], 32)
	if err != nil {
		die(err)
	}

	yhi, err := strconv.ParseFloat(os.Args[4], 32)
	if err != nil {
		die(err)
	}

	npoints, err := strconv.Atoi(os.Args[5])
	if err != nil {
		die(err)
	}

	xint := algs4.NewInterval1D(xlo, xhi)
	yint := algs4.NewInterval1D(ylo, yhi)
	box := algs4.NewInterval2D(xint, yint, nil)
	box.Draw()

	c := algs4.NewCounter("hits")
	for t := 0; t < npoints; t++ {
		x := rand.Float64()
		y := rand.Float64()
		p := algs4.NewPoint2D(x, y, nil)
		if box.Contains(p) {
			c.Increment()
		} else {
			p.Draw()
		}
	}

	algs4.Draw()
	fmt.Println(c)
	fmt.Printf("area = %.2f\n", box.Area())
}
