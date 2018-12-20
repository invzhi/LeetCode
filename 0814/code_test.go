package leetcode

import (
	"reflect"
	"testing"
)

func TestPruneTree(t *testing.T) {
	var tests = []struct {
		before []interface{}
		after  *TreeNode
	}{
		{[]interface{}{1, nil, 0, 0, 1}, newTree(1, nil, 0, nil, 1)},
		{[]interface{}{1, 0, 1, 0, 0, 0, 1}, newTree(1, nil, 1, nil, 1)},
		{[]interface{}{1, 1, 0, 1, 1, 0, 1, 0}, newTree(1, 1, 0, 1, 1, nil, 1)},
	}

	for _, tt := range tests {
		after := pruneTree(newTree(tt.before...))
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("pruneTree(%v) return %v, want %v", newTree(tt.before...), after, tt.after)
		}
	}
}
