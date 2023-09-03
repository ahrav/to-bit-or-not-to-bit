package bitwise_and

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountSetBit(t *testing.T) {
	tests := []struct {
		input    uint64
		expected uint64
	}{
		{0, 0},                     // No set bits
		{1, 1},                     // Single set bit
		{10, 2},                    // Multiple set bits
		{255, 8},                   // All bits set
		{9223372036854775807, 63},  // All bits set except MSB
		{9223372036854775808, 1},   // MSB set
		{18446744073709551615, 64}, // All bits set
		{18446744073709551614, 63}, // All bits set except LSB
		{18446744073709551613, 63}, // All bits set except 2nd LSB
		{18446744073709551612, 62}, // All bits set except 2nd and 3rd LSB
		{18446744073709551611, 63}, // All bits set except 4th LSB
		{18446744073709551610, 62}, // All bits set except 4th and 5th LSB
	}

	for _, test := range tests {
		result := countSetBit(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkCountSetBit(b *testing.B) {
	input := uint64(18446744073709551615)

	for n := 0; n < b.N; n++ {
		countSetBit(input)
	}
}

func TestCountSetBitWithTable(t *testing.T) {
	tests := []struct {
		input    uint64
		expected uint64
	}{
		{0, 0},                     // No set bits
		{1, 1},                     // Single set bit
		{10, 2},                    // Multiple set bits
		{255, 8},                   // All bits set
		{9223372036854775807, 63},  // All bits set except MSB
		{9223372036854775808, 1},   // MSB set
		{18446744073709551615, 64}, // All bits set
		{18446744073709551614, 63}, // All bits set except LSB
		{18446744073709551613, 63}, // All bits set except 2nd LSB
		{18446744073709551612, 62}, // All bits set except 2nd and 3rd LSB
		{18446744073709551611, 63}, // All bits set except 4th LSB
		{18446744073709551610, 62}, // All bits set except 4th and 5th LSB
	}

	for _, test := range tests {
		result := countSetBitWithTable(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkCountSetBitWithTable(b *testing.B) {
	input := uint64(18446744073709551615)

	for n := 0; n < b.N; n++ {
		countSetBitWithTable(input)
	}
}

func TestCountingBits(t *testing.T) {
	tests := []struct {
		n        uint64
		expected []uint64
	}{
		{0, []uint64{0}},
		{1, []uint64{0, 1}},
		{2, []uint64{0, 1, 1}},
		{3, []uint64{0, 1, 1, 2}},
		{4, []uint64{0, 1, 1, 2, 1}},
		{5, []uint64{0, 1, 1, 2, 1, 2}},
		{8, []uint64{0, 1, 1, 2, 1, 2, 2, 3, 1}},
		{16, []uint64{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}},
	}

	for _, test := range tests {
		result := countingBits(test.n)
		if !cmp.Equal(result, test.expected) {
			t.Errorf("countingBits(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n        uint64
		expected bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{234432, true},
		{12345632832, true},
	}

	for _, test := range tests {
		result := isEven(test.n)
		if result != test.expected {
			t.Errorf("isEven(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

var result bool

func BenchmarkIsEven(b *testing.B) {
	input := uint64(12345632832)

	var r bool
	for n := 0; n < b.N; n++ {
		r = isEven(input)
	}
	result = r
}

func BenchmarkIsEvenNaive(b *testing.B) {
	input := uint64(12345632832)

	var r bool
	for n := 0; n < b.N; n++ {
		r = isEvenNaive(input)
	}
	result = r
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n        uint64
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{8, true},
		{16, true},
		{32, true},
		{33, false},
		{68719476736, true},
		{137438953472, true},
		{562949953421312, true},
		{1125899906842624, true},
	}

	for _, test := range tests {
		result := isPowerOfTwo(test.n)
		if result != test.expected {
			t.Errorf("isPowerOfTwo(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func BenchmarkIsPowerOfTwo(b *testing.B) {
	input := uint64(1125899906842624)

	var r bool
	for n := 0; n < b.N; n++ {
		r = isPowerOfTwo(input)
	}
	result = r
}

func TestIsPowerOfTwoNaive(t *testing.T) {
	tests := []struct {
		n        uint64
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{8, true},
		{16, true},
		{32, true},
		{33, false},
		{68719476736, true},
		{137438953472, true},
		{562949953421312, true},
		{1125899906842624, true},
	}

	for _, test := range tests {
		result := isPowerOfTwoNaive(test.n)
		if result != test.expected {
			t.Errorf("isPowerOfTwo(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func BenchmarkIsPowerOfTwoNaive(b *testing.B) {
	input := uint64(1125899906842624)

	var r bool
	for n := 0; n < b.N; n++ {
		r = isPowerOfTwoNaive(input)
	}
	result = r
}
