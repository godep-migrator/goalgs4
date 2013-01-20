package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	. "github.com/meatballhat/goalgs4"
)

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	a := []int64{}

	for i := int64(0); i < n; i++ {
		a = append(a, int64(RandomUniform(float64(-1000000.0), float64(1000000.0))))
	}

	timer := NewStopwatch()
	cnt := count(a)
	elapsed := timer.ElapsedTime()

	fmt.Printf("%v triples %v\n", cnt, elapsed)
}

func count(a []int64) int {
	n := len(a)
	cnt := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (a[i] + a[j] + a[k]) == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}
