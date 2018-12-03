package leetcode

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	var tests = []struct {
		before []interface{}
		key    int
		after  *TreeNode
	}{
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 0, newTree(5, 3, 6, 2, 4, nil, 7)},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 3, newTree(5, 4, 6, 2, nil, nil, 7)},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 4, newTree(5, 3, 6, 2, nil, nil, 7)},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 5, newTree(6, 3, 7, 2, 4)},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 6, newTree(5, 3, 7, 2, 4)},
		{[]interface{}{5, 3, 6, 2, 4, 7}, 6, newTree(5, 3, 7, 2, 4)},
		{[]interface{}{4, 2, 6, 1, 3, 5, 7}, 4, newTree(5, 2, 6, 1, 3, nil, 7)},
	}

	for _, tt := range tests {
		after := deleteNode(newTree(tt.before...), tt.key)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("deleteNode(%v, %v) return %v, want %v", newTree(tt.before...), tt.key, after, tt.after)
		}
	}
}
