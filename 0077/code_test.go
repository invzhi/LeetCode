package leetcode

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	var tests = []struct {
		n, k         int
		combinations [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, tt := range tests {
		combinations := combine(tt.n, tt.k)
		if reflect.DeepEqual(combinations, tt.combinations) == false {
			t.Errorf("combine(%v, %v) return %v, want %v", tt.n, tt.k, combinations, tt.combinations)
		}
	}
}
