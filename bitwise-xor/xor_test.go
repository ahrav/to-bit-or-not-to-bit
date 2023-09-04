package bitwise_xor

import "testing"

func TestSwap(t *testing.T) {
	testCases := []struct {
		a         int32
		b         int32
		expectedA int32
		expectedB int32
	}{
		{1, 2, 2, 1},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{42, 123, 123, 42},
	}

	for _, tc := range testCases {
		a, b := Swap(tc.a, tc.b)
		if a != tc.expectedA || b != tc.expectedB {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", tc.expectedA, tc.expectedB, a, b)
		}
	}
}

var result1, result2 int32

func BenchmarkSwap(b *testing.B) {
	var r1, r2 int32
	for i := 0; i < b.N; i++ {
		r1, r2 = Swap(1, 2)
	}
	result1, result2 = r1, r2
}

func BenchmarkSwapNaive(b *testing.B) {
	var r1, r2 int32
	for i := 0; i < b.N; i++ {
		r1, r2 = SwapNaive(1, 2)
	}
	result1, result2 = r1, r2
}

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 2, 3, 3}, 1},
		{[]int{1, 1, 2, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3, 3, 4}, 4},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 4}, 4},
		{[]int{4, 1, 2, 9, 1, 4, 2}, 9},
	}

	for _, tc := range testCases {
		actual := FindOdd(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func TestFindOddNaive(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 2, 3, 3}, 1},
		{[]int{1, 1, 2, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3, 3, 4}, 4},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 4}, 4},
	}

	for _, tc := range testCases {
		actual := FindOddNaive(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

var result int

func BenchmarkFindOdd(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = FindOdd([]int{1, 2, 2, 3, 3})
	}
	result = r
}

func BenchmarkFindOddNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FindOddNaive([]int{1, 2, 2, 3, 3})
	}
}

func TestHammingDistance(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 8, 2},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{4, 14, 2},
	}

	for _, tc := range testCases {
		actual := HammingDistance(tc.x, tc.y)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func TestHammingDistanceKernighan(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 8, 2},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{4, 14, 2},
	}

	for _, tc := range testCases {
		actual := HamminDistanceKernighan(tc.x, tc.y)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func BenchmarkHammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HammingDistance(12389324, 83282349)
	}
}

func BenchmarkHammingDistanceKernighan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HamminDistanceKernighan(12389324, 83282349)
	}
}

func TestMissingNumberNaive(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 0, 2, 3, 5}, 4},
		{[]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 10}, 9},
		{[]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
	}

	for _, tc := range testCases {
		actual := MissingNumberNaive(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func TestMissingNumberSum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 0, 2, 3, 5}, 4},
		{[]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 10}, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
	}

	for _, tc := range testCases {
		actual := MissingNumberSum(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func TestMissingNumberXOR(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 0, 2, 3, 5}, 4},
		{[]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 10}, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
	}

	for _, tc := range testCases {
		actual := MissingNumberXOR(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, actual)
		}
	}
}

func consecutiveArrMissingOne(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return append(arr[:n/2], arr[n/2+1:]...)
}

func BenchmarkMissingNumberNaive(b *testing.B) {
	arr := consecutiveArrMissingOne(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MissingNumberNaive(arr)
	}
}

func BenchmarkMissingNumberSum(b *testing.B) {
	var r int
	arr := consecutiveArrMissingOne(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = MissingNumberSum(arr)
	}
	result = r
}

func BenchmarkMissingNumberXOR(b *testing.B) {
	var r int
	arr := consecutiveArrMissingOne(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = MissingNumberXOR(arr)
	}
	result = r
}
