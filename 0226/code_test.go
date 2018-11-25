package leetcode

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	var tests = []struct {
		before []interface{}
		after  *TreeNode
	}{
		{[]interface{}{4, 2, 7, 1, 3, 6, 9}, newTree(4, 7, 2, 9, 6, 3, 1)},
	}

	for _, tt := range tests {
		after := invertTree(newTree(tt.before...))
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("invertTree(%v) return %v, want %v", newTree(tt.before), after, tt.after)
		}
	}
}
