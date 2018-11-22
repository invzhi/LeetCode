package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	var tests = []struct {
		t1, t2 *TreeNode
		merged *TreeNode
	}{
		{newTree(1, 3, 2, 5), newTree(2, 1, 3, nil, 4, nil, 7), newTree(3, 4, 5, 5, 4, nil, 7)},
	}

	for _, tt := range tests {
		merged := mergeTrees(tt.t1, tt.t2)
		if reflect.DeepEqual(merged, tt.merged) == false {
			t.Errorf("mergeTrees(%v, %v) return %v, want %v", tt.t1, tt.t2, merged, tt.merged)
		}
	}
}
