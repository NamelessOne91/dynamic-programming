package fibonacci

import "testing"

func TestFibonacciNumber(t *testing.T) {
	tt := []struct {
		input    int
		expected int
	}{{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {10, 55}, {21, 10946}, {50, 12586269025}}

	for _, test := range tt {
		n := FibonacciNumber(test.input)
		if test.expected != n {
			t.Fatalf("Expected %d - got %d", test.expected, n)
		}
	}
}
