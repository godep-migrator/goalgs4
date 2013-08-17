package main

import (
	"fmt"
	"math"
	"os"

	. "goalgs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	var numbers Bag
	if len(os.Args) > 1 && os.Args[1] == "ll" {
		numbers = NewLinkedListBag()
	} else {
		numbers = NewArrayBag()
	}

	for !Stdin.IsEmpty() {
		d, err := Stdin.ReadDouble()
		if err != nil {
			break
		}
		numbers.Add(d)
	}

	n := numbers.Size()
	sum := float64(0.0)

	for cursor := numbers.First(); cursor != nil; cursor = cursor.Next() {
		sum += cursor.Value().Float()
	}

	mean := sum / float64(n)

	sum = float64(0.0)
	for cursor := numbers.First(); cursor != nil; cursor = cursor.Next() {
		val := cursor.Value().Float()
		sum += (val - mean) * (val - mean)
	}

	std := math.Sqrt(sum / (float64(n) - float64(1)))

	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Std dev: %.2f\n", std)
}
