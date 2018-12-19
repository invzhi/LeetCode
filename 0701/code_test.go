package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertIntoBST(t *testing.T) {
	var tests = []struct {
		before []interface{}
		val    int
		after  *TreeNode
	}{
		{[]interface{}{4, 2, 7, 1, 3}, 5, newTree(4, 2, 7, 1, 3, 5)},
		{[]interface{}{40, 20, 60, 10, 30, 50, 70}, 25, newTree(40, 20, 60, 10, 30, 50, 70, nil, nil, 25)},
		{[]interface{}{40, 20, 60, 10, 30, 50, 70}, 35, newTree(40, 20, 60, 10, 30, 50, 70, nil, nil, nil, 35)},
		{[]interface{}{}, 5, newTree(5)},
	}

	for _, tt := range tests {
		after := insertIntoBST(newTree(tt.before...), tt.val)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("insertIntoBST(%v, %v) return %v, want %v", newTree(tt.before...), tt.val, after, tt.after)
		}
	}
}
