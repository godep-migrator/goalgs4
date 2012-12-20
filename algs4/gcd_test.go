package algs4_test

import (
	"testing"
)

import (
	"github.com/meatballhat/box-o-sand/algs4/src/go/algs4"
)

func TestGcd(t *testing.T) {
	d := algs4.Gcd(uint64(81), uint64(72))
	if d != 9 {
		t.Error("WRONG")
	}
}
