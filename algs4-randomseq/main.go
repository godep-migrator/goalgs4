package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

import (
	"github.com/meatballhat/algs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 4 {
		die(errors.New("We need an int and two floats!"))
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		die(err)
	}

	lo, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		die(err)
	}

	hi, err := strconv.ParseFloat(os.Args[3], 32)
	if err != nil {
		die(err)
	}

	for i := 0; i < n; i++ {
		x := algs4.RandomUniform(lo, hi)
		fmt.Printf("%.2f\n", x)
	}
}
