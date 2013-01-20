package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/meatballhat/goalgs4"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	a, err := ReadInts(fd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count(a))
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
