package leetcode

import (
	"reflect"
	"testing"
)

func TestLargestValues(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		vals []int
	}{
		{newTree(1, 3, 2, 5, 3, nil, 9), []int{1, 3, 9}},
	}

	for _, tt := range tests {
		vals := largestValues(tt.tree)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("largestValues(%v) return %v, want %v", tt.tree, vals, tt.vals)
		}
	}
}
