package bitwise_or

import "testing"

func TestMinFlips(t *testing.T) {
	testCases := []struct {
		a, b, c  int32
		expected int32
	}{
		{2, 6, 5, 3},
		{5, 7, 3, 2},
		{0, 0, 1, 1},
	}

	for _, tc := range testCases {
		actual := MinFlips(tc.a, tc.b, tc.c)
		if actual != tc.expected {
			t.Errorf("MinFlips(%d, %d, %d) = %d, want %d",
				tc.a, tc.b, tc.c, actual, tc.expected)
		}
	}

}

func BenchmarkMinFlips(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinFlips(2, 6, 5)
	}
}
