package bitwise_and

// CountSetBit using Brian Kernighan's algorithm.
func CountSetBit(x uint64) uint64 {
	var count uint64
	for x != 0 {
		x = x & (x - 1)
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

// CountSetBitWithTable using a lookup table.
func CountSetBitWithTable(x uint64) uint64 {
	var count uint64
	for i := 0; i < 8; i++ {
		count += lookupTable[x&0xff]
		x >>= 8
	}

	return count
}
