package change

import "testing"

func TestMinimumCoinsNum(t *testing.T) {
	tt := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{"0 amount", []int{}, 0, 0},
		{"straightforward solution", []int{1, 10, 15}, 42, 5},
		{"should not use highest coin", []int{1, 4, 5}, 13, 3},
		{"higher amount value", []int{1, 3, 5, 7, 10}, 114, 12},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			n := MinimumCoinsNum(test.coins, test.amount)
			if test.expected != n {
				t.Fatalf("Expected minimum number of coins to be %d - got %d", test.expected, n)
			}
		})
	}
}

func TestPossibleCoinsCombinations(t *testing.T) {
	tt := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{"0 amount", []int{}, 0, 1},
		{"simple", []int{1, 2}, 2, 2},
		{"order matters (1-4 || 4-1)", []int{1, 4, 5}, 5, 4},
		{"no combinations", []int{2}, 3, 0},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			n := PossibleCoinsCombinations(test.coins, test.amount)
			if test.expected != n {
				t.Fatalf("Expected number of possible combinations to be %d - got %d", test.expected, n)
			}
		})
	}
}
