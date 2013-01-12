package algs4_test

import (
	"testing"
)

import (
	"github.com/meatballhat/goalgs4"
)

func TestBinarySearchRank(t *testing.T) {
	arr := [...]int{0, 2, 9, 14, 23}

	if algs4.BinarySearchRank(9, arr[:]) != 2 {
		t.Error("Failed to find 9 in", arr)
	}

	if algs4.BinarySearchRank(0, arr[:]) != 0 {
		t.Error("Failed to find 0 in", arr)
	}
}
