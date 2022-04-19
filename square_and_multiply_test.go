package squareandmultiply_test

import (
	"fmt"
	"testing"

	squareandmultiply "github.com/dangermike/square_and_multiply"
)

func TestABmodC(t *testing.T) {
	for _, test := range []struct {
		base, exp, mod, expected uint
	}{
		{0, 123456, 5, 0},
		{1, 123456, 5, 1},
		{123456, 0, 543, 1},
		{123456, 123456, 1, 0},
		{3, 45, 7, 6},
		{23, 373, 747, 131},
		{13, 1 << 63, 65534, 29023},
	} {
		t.Run(fmt.Sprintf("%d**%d mod %d", test.base, test.exp, test.mod), func(t *testing.T) {
			actual := squareandmultiply.ABmodC(test.base, test.exp, test.mod)
			if test.expected != actual {
				t.Fatalf(
					"\nexpected: %d\n  actual: %d",
					test.expected, actual,
				)
			}
		})
	}
}

func BenchmarkABmodC(b *testing.B) {
	for _, test := range []struct {
		base, exp, mod uint
	}{
		{2, 0xFFFFFFFFFFFFFFFF, 876543},
		{2, 0xFFFFFFFFFFFFFFFF, 4},
		{2, 0x7000000000000000, 876543},
		{2, (1 << 32) + 1, 876543},
		{2, 765432, 876543},
		{2, 65567, 876543},
		{2, 262145, 876543},
		{2, 1, 876543},
	} {
		b.Run(fmt.Sprintf("%d**%d mod %d", test.base, test.exp, test.mod), func(b *testing.B) {
			var base, exp, mod uint = test.base, test.exp, test.mod
			for i := 0; i < b.N; i++ {
				_ = squareandmultiply.ABmodC(base, exp, mod)
			}
		})
	}
}
