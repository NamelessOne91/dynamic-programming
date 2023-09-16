package change

import "slices"

// MinimumCoinsNum returns the minimum number of the given coins necessary to produce the given amount
//
// Bottom-up implementation
func MinimumCoinsNum(coins []int, amount int) int {
	memo := map[int]int{0: 0}

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			subProblem := i - coin
			if subProblem < 0 {
				continue
			}

			values := []int{memo[subProblem] + 1}
			v, ok := memo[i]
			if ok {
				values = append(values, v)
			}

			memo[i] = slices.Min(values)
		}
	}

	return memo[amount]
}

// PossibleCoinsCombinations returns the number of ways the given coins can be combined to form the given amount (order matters)
//
// Bottom-up implementation
func PossibleCoinsCombinations(coins []int, amount int) int {
	memo := map[int]int{0: 1}

	for i := 1; i < amount+1; i++ {
		memo[i] = 0

		for _, coin := range coins {
			subProblem := i - coin
			if subProblem < 0 {
				continue
			}

			memo[i] = memo[i] + memo[subProblem]
		}
	}

	return memo[amount]
}
