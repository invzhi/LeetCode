package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindFrequentTreeSum(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		sums []int
	}{
		{newTree(5, 2, -3), []int{-3, 2, 4}},
		{newTree(5, 2, -5), []int{2}},
	}

	for _, tt := range tests {
		sums := findFrequentTreeSum(tt.tree)
		sort.Ints(sums)
		if reflect.DeepEqual(sums, tt.sums) == false {
			t.Errorf("findFrequentTreeSum(%v) return %v, want %v", tt.tree, sums, tt.sums)
		}
	}
}
