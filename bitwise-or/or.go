package bitwise_or

// MinFlips returns the minimum number of bit flips required to make a OR b equal to c.
func MinFlips(a, b, c int32) int32 {
	var flips int32
	for i := 0; i < 32; i++ {
		bitC := (c >> i) & 1
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1

		if (bitA | bitB) != bitC {
			if bitC == 0 {
				flips += bitA + bitB // This will add 2 if both A and B bits are 1, otherwise 1.
			} else {
				flips++
			}
		}

	}

	return flips
}
