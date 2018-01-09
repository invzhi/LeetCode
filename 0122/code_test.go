package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		profit int
	}{
		{[]int{4, 6, 7, 2, 8}, 9},
		{[]int{6, 5, 4, 3, 2}, 0},
		{[]int{4, 2, 6, 5, 8}, 7},
		{[]int{1, 9}, 8},
		{[]int{1}, 0},
		{nil, 0},
	}

	for _, tt := range tests {
		profit := maxProfit(tt.prices)
		if profit != tt.profit {
			t.Errorf("maxProfit(%v) return %v, want %v", tt.prices, profit, tt.profit)
		}
	}
}
