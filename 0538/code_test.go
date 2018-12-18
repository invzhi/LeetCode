package leetcode

import (
	"reflect"
	"testing"
)

func TestConvertBST(t *testing.T) {
	var tests = []struct {
		before []interface{}
		after  *TreeNode
	}{
		{[]interface{}{5, 2, 13}, newTree(18, 20, 13)},
	}

	for _, tt := range tests {
		after := convertBST(newTree(tt.before...))
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("convertBST(%v) return %v, want %v", newTree(tt.before...), after, tt.after)
		}
	}
}
