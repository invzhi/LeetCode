package leetcode

import (
	"reflect"
	"testing"
)

func findTarget(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}

	if t := findTarget(root.Left, target); t != nil {
		return t
	}
	if t := findTarget(root.Right, target); t != nil {
		return t
	}
	return nil
}

func TestDistanceK(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		target int
		k      int
		vals   []int
	}{
		{newTree(0, 1, nil, 3, 2), 2, 1, []int{1}},
		{newTree(0, 1, nil, 3, 2), 2, 2, []int{3, 0}},
		{newTree(0, nil, 1, 2, 3), 2, 1, []int{1}},
		{newTree(0, nil, 1, 2, 3), 2, 2, []int{3, 0}},
		{newTree(0, 1, nil, 3, 2, 4), 2, 1, []int{1}},
		{newTree(0, 2, 1, nil, nil, 3), 3, 3, []int{2}},
		{newTree(0, 1, nil, nil, 2, nil, 3, nil, 4), 3, 0, []int{3}},
		{newTree(0, nil, 1, 2, 3, nil, nil, nil, 4), 2, 1, []int{1}},
		{newTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4), 5, 2, []int{7, 4, 1}},
	}

	for _, tt := range tests {
		target := findTarget(tt.tree, tt.target)
		vals := distanceK(tt.tree, target, tt.k)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("distanceK(%v, %v, %v) return %v, want %v", tt.tree, tt.target, tt.k, vals, tt.vals)
		}
	}
}
