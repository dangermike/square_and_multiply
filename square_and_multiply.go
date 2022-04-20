package squareandmultiply

import (
	"math/bits"
)

func ABmodC(base, exp, mod uint) uint {
	// mods can't be zero
	if mod == 0 {
		panic("mod value cannot be zero")
	}
	if mod == 1 {
		return 0
	}
	if base == 1 || exp == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}

	// Starting with base rather than 1, we have effectively consumed the
	// first of the bits
	acc := base % mod

	// +1 to discard the first bit
	lz := uint(bits.LeadingZeros(exp)) + 1

	// reverse and left-shift so we can consume bits as an instruction stream
	ops := bits.Reverse(exp) >> lz

	// 64 - lz is the number of set bits, skipping the first
	for i := uint(0); i < 64-lz; i++ {
		v := ops & 1
		ops >>= 1
		acc = (acc * acc) % mod
		if v == 1 {
			acc = (acc * base) % mod
		}
	}

	return acc
}
