package squareandmultiply_test

import (
	"fmt"
	"math/big"
	"testing"

	squareandmultiply "github.com/dangermike/square_and_multiply"
)

func bigABmodC(base, exp, mod uint) uint {
	return uint(big.NewInt(0).Exp(
		big.NewInt(0).SetUint64(uint64(base)),
		big.NewInt(0).SetUint64(uint64(exp)),
		big.NewInt(0).SetUint64(uint64(mod)),
	).Int64())
}

func DoTest(t *testing.T, base uint, exp uint, mod uint) {
	if mod == 0 {
		t.SkipNow()
	}

	// using the big.Int.Exp implementation as a reference
	expected := bigABmodC(base, exp, mod)
	actual := squareandmultiply.ABmodC(base, exp, mod)

	if expected != actual {
		t.Fatalf(
			"\n    expr: (%d ** %d) mod %d"+
				"\nexpected: %d"+
				"\n  actual: %d",
			base, exp, mod, expected, actual,
		)
	}
}

func TestABmodC(t *testing.T) {
	for _, test := range []struct {
		base, exp, mod uint
	}{
		{0, 123456, 5},
		{1, 123456, 5},
		{123456, 0, 543},
		{123456, 123456, 1},
		{3, 45, 7},
		{23, 373, 747},
		{13, 1 << 63, 65534},
	} {
		t.Run(fmt.Sprintf("%d**%d mod %d", test.base, test.exp, test.mod), func(t *testing.T) {
			DoTest(t, test.base, test.exp, test.mod)
		})
	}
}

func FuzzABmodC(f *testing.F) {
	f.Add(uint(1), uint(1), uint(1))
	f.Fuzz(DoTest)
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
