package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	var tests = []struct {
		candidates   []int
		target       int
		combinations [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 2, 5}, {1, 7}, {1, 1, 6}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for _, tt := range tests {
		combinations := combinationSum2(tt.candidates, tt.target)
		if reflect.DeepEqual(combinations, tt.combinations) == false {
			t.Errorf("combinationSum2(%v, %v) return %v, want %v", tt.candidates, tt.target, combinations, tt.combinations)
		}
	}
}
