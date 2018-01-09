package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		profit int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{nil, 0},
	}

	for _, tt := range tests {
		profit := maxProfit(tt.prices)
		if profit != tt.profit {
			t.Errorf("maxProfit(%v) return %v, want %v", tt.prices, profit, tt.profit)
		}
	}
}
