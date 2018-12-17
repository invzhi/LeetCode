package leetcode

import (
	"reflect"
	"testing"
)

func TestAddOneRow(t *testing.T) {
	var tests = []struct {
		before []interface{}
		v, d   int
		after  *TreeNode
	}{
		{[]interface{}{4, 2, 6, 3, 1, 5}, 1, 1, newTree(1, 4, nil, 2, 6, 3, 1, 5)},
		{[]interface{}{4, 2, 6, 3, 1, 5}, 1, 2, newTree(4, 1, 1, 2, nil, nil, 6, 3, 1, 5)},
		{[]interface{}{4, 2, nil, 3, 1}, 1, 3, newTree(4, 2, nil, 1, 1, 3, nil, nil, 1)},
	}

	for _, tt := range tests {
		after := addOneRow(newTree(tt.before...), tt.v, tt.d)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("addOneRow(%v, %v, %v) return %v, want %v", newTree(tt.before...), tt.v, tt.d, after, tt.after)
		}
	}
}
