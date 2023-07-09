package bitwise_and

import "math"

// Find the number of set bits for n using Brian Kernighan's algorithm.
// If n is < 0 or > 64 bits, return 0.
func countSetBit(n uint64) uint64 {
	// Validate input isn't negative or larger than 64 bits.
	if n < 0 || n > math.MaxUint64 {
		return 0
	}

	var count uint64
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

var lookupTable [256]uint64

func init() {
	lookupTable[0] = 0
	for i := 1; i < 256; i++ {
		lookupTable[i] = uint64(i&1) + lookupTable[i>>1]
	}
}

// Find the number of set bits for n using a lookup table.
// If n is < 0 or > 64 bits, return 0.
func countSetBitWithTable(n uint64) uint64 {
	// Validate input isn't negative or larger than 64 bits.
	if n < 0 || n > math.MaxUint64 {
		return 0
	}

	var count uint64
	for i := 0; i < 8; i++ {
		count += lookupTable[n&0xff]
		n >>= 8
	}

	return count
}

// Finds the number of 1's for every number from 0 to n.
func countingBits(n uint64) []uint64 {
	result := make([]uint64, n+1)
	for i := uint64(0); i <= n; i++ {
		result[i] = countSetBit(i)
	}
	return result
}
