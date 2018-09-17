package leetcode

import (
	"reflect"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	var tests = []struct {
		a, b     []int
		exchange []int
	}{
		{[]int{1, 1}, []int{2, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{2, 3}, []int{1, 2}},
		{[]int{2}, []int{1, 3}, []int{2, 3}},
		{[]int{1, 2, 5}, []int{2, 4}, []int{5, 4}},
		{nil, nil, nil},
	}

	for _, tt := range tests {
		exchange := fairCandySwap(tt.a, tt.b)
		if reflect.DeepEqual(exchange, tt.exchange) == false {
			t.Errorf("fairCandySwap(%v, %v) return %v, want %v", tt.a, tt.b, exchange, tt.exchange)
		}
	}
}
