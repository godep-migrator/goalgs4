package main

import (
	"fmt"
	"math"
	"os"

	"github.com/meatballhat/goalgs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	numbers := algs4.NewArrayBag()

	for !algs4.Stdin.IsEmpty() {
		d, err := algs4.Stdin.ReadDouble()
		if err != nil {
			break
		}
		numbers.Add(d)
	}

	n := numbers.Size()
	sum := float64(0.0)

	for cursor := numbers.First(); cursor != nil; cursor = cursor.Next() {
		v := cursor.Value()
		if val, ok := v.(float64); ok {
			sum += val
		} else {
			fmt.Fprintf(os.Stderr, "%v is not a float64!", v)
		}
	}

	mean := sum / float64(n)

	sum = float64(0.0)
	for cursor := numbers.First(); cursor != nil; cursor = cursor.Next() {
		v := cursor.Value()
		if val, ok := v.(float64); ok {
			sum += (val - mean) * (val - mean)
		} else {
			fmt.Fprintf(os.Stderr, "%v is not a float64!", v)
		}
	}

	std := math.Sqrt(sum / (float64(n) - float64(1)))

	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Std dev: %.2f\n", std)
}
