package bitwise_and

import "testing"

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
		result := CountSetBit(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkCountSetBit(b *testing.B) {
	// Set up a test case or input for the benchmark
	input := uint64(18446744073709551615)

	// Run the benchmark
	for n := 0; n < b.N; n++ {
		CountSetBit(input)
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
		result := CountSetBitWithTable(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkCountSetBitWithTable(b *testing.B) {
	// Set up a test case or input for the benchmark
	input := uint64(18446744073709551615)

	// Run the benchmark
	for n := 0; n < b.N; n++ {
		CountSetBitWithTable(input)
	}
}
