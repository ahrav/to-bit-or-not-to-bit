package bitwise_not

import "testing"

func TestSignFlip(t *testing.T) {
	testCases := []struct {
		n        int32
		expected int32
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{2, -2},
		{42, -42},
	}

	for _, tc := range testCases {
		actual := SignFlip(tc.n)
		if actual != tc.expected {
			t.Errorf("SignFlip(%d) = %d, want %d",
				tc.n, actual, tc.expected)
		}
	}
}

var result int32

func BenchmarkSignFlip(b *testing.B) {
	var r int32
	for i := 0; i < b.N; i++ {
		r = SignFlip(18492349)
	}
	result = r
}
