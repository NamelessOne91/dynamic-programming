package fibonacci

// FibonacciNumber return the n-th number in the Fibonacci sequence
//
// Bottom-up implementation
func FibonacciNumber(n int) int {
	memo := make(map[int]int)

	result := 1
	for i := 1; i < n+1; i++ {
		if i > 2 {
			result = memo[i-1] + memo[i-2]
		}
		memo[i] = result
	}

	return memo[n]
}
