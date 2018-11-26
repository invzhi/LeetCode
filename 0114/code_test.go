package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	var tests = []struct {
		before []interface{}
		after  *TreeNode
	}{
		{[]interface{}{1, 2, 5, 3, 4, nil, 6}, newTree(1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6)},
	}

	for _, tt := range tests {
		after := newTree(tt.before...)
		flatten(after)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("flatten(%v) return %v, want %v", newTree(tt.before...), after, tt.after)
		}
	}
}
