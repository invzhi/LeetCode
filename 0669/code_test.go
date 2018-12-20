package leetcode

import (
	"reflect"
	"testing"
)

func TestTrimBST(t *testing.T) {
	var tests = []struct {
		before []interface{}
		L, R   int
		after  *TreeNode
	}{
		{[]interface{}{1, 0, 2}, 1, 2, newTree(1, nil, 2)},
		{[]interface{}{3, 0, 4, nil, 2, nil, nil, 1}, 1, 3, newTree(3, 2, nil, 1)},
	}

	for _, tt := range tests {
		after := trimBST(newTree(tt.before...), tt.L, tt.R)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("trimBST(%v, %v, %v) return %v, want %v", newTree(tt.before...), tt.L, tt.R, after, tt.after)
		}
	}
}
