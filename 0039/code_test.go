package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var tests = []struct {
		candidates   []int
		target       int
		combinations [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, tt := range tests {
		combinations := combinationSum(tt.candidates, tt.target)
		if reflect.DeepEqual(combinations, tt.combinations) == false {
			t.Errorf("combinationSum(%v, %v) return %v, want %v", tt.candidates, tt.target, combinations, tt.combinations)
		}
	}
}
