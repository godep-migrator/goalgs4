package main

import (
	"fmt"
	"io"
	"os"
)

import (
	"github.com/meatballhat/algs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	sum := 0.0
	cnt := 0

	for {
		d, err := algs4.Stdin.ReadDouble()
		if err == io.EOF {
			break
		}

		if err != nil {
			die(err)
		}

		sum += d
		cnt++
	}

	fmt.Printf("Average is %.5f\n", sum/float64(cnt))
}
