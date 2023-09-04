package bitwise_xor

// go:noinline
func Swap(a, b int32) (int32, int32) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// Go compiler is smart enough to optimize this.
// ARM using g c:1.21
// MOVW    main.b+4(FP), R0
// MOVW    R0, main.~r0+8(FP)
// MOVW    main.a(FP), R0
// MOVW    R0, main.~r1+12(FP)
// JMP     (R14)
func SwapNaive(a, b int32) (int32, int32) {
	var tmp int32
	tmp = a
	a = b
	b = tmp
	return a, b
}

// go:noinline
func FindOdd(arr []int) int {
	var result int
	for _, v := range arr {
		result ^= v
	}
	return result
}

func FindOddNaive(arr []int) int {
	var result int
	seen := make(map[int]int)
	for _, v := range arr {
		seen[v]++
	}
	for k, v := range seen {
		if v%2 == 1 {
			result = k
			break
		}
	}
	return result
}

// go:noinline
func HammingDistance(x, y int) int {
	var result int
	xor := x ^ y
	for xor > 0 {
		result += xor & 1
		xor >>= 1
	}
	return result
}

// go:noinline
func HamminDistanceKernighan(x, y int) int {
	var result int
	xor := x ^ y
	for xor > 0 {
		result++
		xor &= xor - 1
	}
	return result
}

func MissingNumberNaive(arr []int) int {
	// Create a slice where its the size of the array
	// and set all values to -1.
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	for i := 0; i < len(arr); i++ {
		if !seen[i] {
			return i
		}
	}
	return -1
}

func MissingNumberSum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return (len(arr) * (len(arr) + 1) / 2) - sum
}

func MissingNumberXOR(arr []int) int {
	xor := len(arr)
	for i, v := range arr {
		xor ^= i ^ v
	}
	return xor
}
