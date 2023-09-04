package bitwise_not

// go:noinline
func SignFlip(n int32) int32 {
	return ^n + 1
}
